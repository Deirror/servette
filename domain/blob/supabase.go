package blob

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/Deirror/servette/transport/http/client"
)

type SupabaseStorageClient struct {
	ProjectURL string
	APIKey     string
	Client     *http.Client
}

func NewSupabaseStorageClient(cfg *Config) (*SupabaseStorageClient, error) {
	defaultClientCfg := client.DefaultConfig()
	return &SupabaseStorageClient{
		ProjectURL: cfg.ProjectURL,
		APIKey:     cfg.APIKey,
		Client:     client.New(&defaultClientCfg),
	}, nil
}

func (s *SupabaseStorageClient) Upload(ctx context.Context, bucket, path string, data []byte) error {
	url := fmt.Sprintf("%s/storage/v1/object/%s/%s", s.ProjectURL, bucket, path)
	req, err := http.NewRequest("PUT", url, bytes.NewReader(data))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.APIKey))
	req.Header.Set("Content-Type", "application/octet-stream")

	resp, err := s.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("failed to upload file: status code %d", resp.StatusCode)
	}

	return nil
}

func (s *SupabaseStorageClient) Download(ctx context.Context, bucket, path string) ([]byte, error) {
	url := fmt.Sprintf("%s/storage/v1/object/%s/%s", s.ProjectURL, bucket, path)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.APIKey))

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to download file: %d - %s", resp.StatusCode, string(body))
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	return data, nil
}

// CreateSignedURL automatically deletes existing object on 409 conflict when forUpload=true
func (s *SupabaseStorageClient) CreateSignedURL(ctx context.Context, bucket, objectPath string, expiresInSec int, forUpload, upsert bool) (string, string, error) {
	signedURL, token, err := s.createSignedURLInternal(ctx, bucket, objectPath, expiresInSec, forUpload, upsert)
	if err != nil && forUpload {
		// check if it’s a 409 Duplicate error
		if strings.Contains(err.Error(), "409") {
			_ = s.deleteObject(ctx, bucket, objectPath) // ignore delete errors
			return s.createSignedURLInternal(ctx, bucket, objectPath, expiresInSec, forUpload, upsert)
		}
	}
	return signedURL, token, err
}

// internal helper: actually sends the POST request
func (s *SupabaseStorageClient) createSignedURLInternal(ctx context.Context, bucket, objectPath string, expiresInSec int, forUpload, upsert bool) (string, string, error) {
	escBucket := url.PathEscape(bucket)
	segs := strings.Split(objectPath, "/")
	for i := range segs {
		segs[i] = url.PathEscape(segs[i])
	}
	escPath := strings.Join(segs, "/")

	var endpoint string
	if forUpload {
		endpoint = fmt.Sprintf("%s/storage/v1/object/upload/sign/%s/%s", s.ProjectURL, escBucket, escPath)
	} else {
		endpoint = fmt.Sprintf("%s/storage/v1/object/sign/%s/%s", s.ProjectURL, escBucket, escPath)
	}

	payload := map[string]interface{}{"expiresIn": expiresInSec}
	if forUpload {
		payload["upsert"] = upsert
	}
	b, _ := json.Marshal(payload)

	req, _ := http.NewRequestWithContext(ctx, "POST", endpoint, bytes.NewReader(b))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.APIKey))
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.Client.Do(req)
	if err != nil {
		return "", "", fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return "", "", fmt.Errorf("create signed url failed: %d - %s", resp.StatusCode, string(body))
	}

	signedURL, token, perr := parseSignedResponse(body)
	if perr != nil {
		return "", "", fmt.Errorf("parse signed response: %w - body: %s", perr, string(body))
	}
	return s.ProjectURL + "/storage/v1" + signedURL, token, nil
}

// deleteObject deletes an object in Supabase Storage (private helper)
func (s *SupabaseStorageClient) deleteObject(ctx context.Context, bucket, objectPath string) error {
	escBucket := url.PathEscape(bucket)
	segs := strings.Split(objectPath, "/")
	for i := range segs {
		segs[i] = url.PathEscape(segs[i])
	}
	escPath := strings.Join(segs, "/")

	endpoint := fmt.Sprintf("%s/storage/v1/object/%s/%s", s.ProjectURL, escBucket, escPath)
	req, _ := http.NewRequestWithContext(ctx, "DELETE", endpoint, nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.APIKey))

	resp, err := s.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusNotFound {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("delete object failed: %d - %s", resp.StatusCode, string(body))
	}
	return nil
}

// parseSignedResponse looks for common keys and returns either a url or a token
func parseSignedResponse(body []byte) (string, string, error) {
	var top map[string]interface{}
	if err := json.Unmarshal(body, &top); err != nil {
		return "", "", fmt.Errorf("invalid json: %w", err)
	}

	tryKeys := func(m map[string]interface{}, keys ...string) (string, bool) {
		for _, k := range keys {
			if v, ok := m[k]; ok {
				if sstr, ok := v.(string); ok && sstr != "" {
					return sstr, true
				}
			}
		}
		return "", false
	}

	// look top-level
	if v, ok := tryKeys(top, "signedURL", "signed_url", "signedUrl", "url"); ok {
		return v, "", nil
	}
	if v, ok := tryKeys(top, "token", "uploadToken"); ok {
		return "", v, nil
	}

	// check data wrapper
	if dataRaw, ok := top["data"]; ok {
		if dataMap, ok := dataRaw.(map[string]interface{}); ok {
			if v, ok := tryKeys(dataMap, "signedURL", "signed_url", "signedUrl", "url"); ok {
				return v, "", nil
			}
			if v, ok := tryKeys(dataMap, "token", "uploadToken"); ok {
				return "", v, nil
			}
		}
	}

	// fallback: scan any string value (heuristic)
	for _, v := range top {
		if sstr, ok := v.(string); ok && sstr != "" {
			if len(sstr) > 4 && (sstr[:4] == "http") {
				return sstr, "", nil
			}
			return "", sstr, nil
		}
	}

	return "", "", fmt.Errorf("no signed url/token in response")
}
