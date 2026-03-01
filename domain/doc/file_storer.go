package doc

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"reflect"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type FileStoreClient struct {
	basePath string
}

func NewFileStoreClient(cfg *Config) *FileStoreClient {
	return &FileStoreClient{
		basePath: cfg.StoreURL,
	}
}

func (f *FileStoreClient) loadCollection(coll string) ([]map[string]any, error) {
	filePath := filepath.Join(f.basePath, coll+".json")

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var rawDocs []map[string]any
	if err := json.Unmarshal(data, &rawDocs); err != nil {
		return nil, err
	}

	return rawDocs, nil
}

func (f *FileStoreClient) FindWithOpts(
	ctx context.Context,
	coll string,
	filter any,
	results any,
	opts *FindOptions,
) error {

	val := reflect.ValueOf(results)
	if val.Kind() != reflect.Ptr {
		return errors.New("results must be pointer")
	}

	rawDocs, err := f.loadCollection(coll)
	if err != nil {
		return err
	}

	filtered := applyFilter(rawDocs, filter)

	// Pagination
	if opts != nil {
		start := opts.Skip
		if start > int64(len(filtered)) {
			filtered = []map[string]any{}
		} else {
			end := start + opts.Limit

			if opts.Limit == 0 {
				end = int64(len(filtered))
			}

			if end > int64(len(filtered)) {
				end = int64(len(filtered))
			}

			filtered = filtered[start:end]
		}
	}

	finalJSON, err := json.Marshal(filtered)
	if err != nil {
		return err
	}

	return json.Unmarshal(finalJSON, results)
}

func (f *FileStoreClient) Find(
	ctx context.Context,
	coll string,
	filter any,
	results any,
) error {
	return f.FindWithOpts(ctx, coll, filter, results, nil)
}

func applyFilter(
	docs []map[string]interface{},
	filter interface{},
) []map[string]interface{} {

	if filter == nil {
		return docs
	}

	f, ok := filter.(bson.M)
	if !ok || len(f) == 0 {
		return docs
	}

	var result []map[string]interface{}

	for _, doc := range docs {
		if matches(doc, f) {
			result = append(result, doc)
		}
	}

	return result
}

func matches(
	doc map[string]interface{},
	filter bson.M,
) bool {

	for key, val := range filter {

		docVal, exists := doc[key]
		if !exists {
			return false
		}

		if !reflect.DeepEqual(docVal, val) {
			return false
		}
	}

	return true
}

func (f *FileStoreClient) Count(
	ctx context.Context,
	coll string,
	filter any,
) (int64, error) {

	rawDocs, err := f.loadCollection(coll)
	if err != nil {
		return 0, err
	}

	filtered := applyFilter(rawDocs, filter)

	return int64(len(filtered)), nil
}

func (f *FileStoreClient) Insert(
	ctx context.Context,
	coll string,
	docs ...any,
) (any, error) {
	return nil, errors.New("not supported")
}

func (f *FileStoreClient) Update(
	ctx context.Context,
	coll string,
	filter any,
	update any,
) (int64, error) {
	return 0, errors.New("not supported")
}

func (f *FileStoreClient) Delete(
	ctx context.Context,
	coll string,
	filter any,
) (int64, error) {
	return 0, errors.New("not supported")
}
