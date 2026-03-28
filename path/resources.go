// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package pathx

import (
	"fmt"
	"net/url"
	"strings"
)

type ResourceKind string

const (
	FilePath ResourceKind = "filepath" // resource from env is file
	URI      ResourceKind = "uri"      // resource from env is uri
	Unknown  ResourceKind = "unknown"  // invalid resource
)

func (r ResourceKind) IsFilePath() bool {
	return r == FilePath
}

func (r ResourceKind) IsURI() bool {
	return r == URI
}

func (r ResourceKind) IsUnknown() bool {
	return r == Unknown
}

// ParseResourceKind converts a string to a ResourceKind. Case-insensitive.
func ParseResourceKind(s string) (ResourceKind, error) {
	normalized := strings.ToLower(strings.TrimSpace(s))
	switch normalized {
	case string(FilePath):
		return FilePath, nil
	case string(URI):
		return URI, nil
	default:
		return "", fmt.Errorf("invalid ResourceKind: %q", s)
	}
}

type Resource string

// Kind checks what is the resource's kind.
func (p Resource) Kind() ResourceKind {
	s := strings.TrimSpace(string(p))
	if s == "" {
		return Unknown
	}

	u, err := url.Parse(s)
	if err == nil && u.Scheme != "" && u.Host != "" {
		return URI
	}

	return FilePath
}

// ResourcesToStrings converts a slice of Resource to a slice of strings.
// Trims whitespace and ignores empty strings.
func ResourcesToStrings(resources []Resource) []string {
	result := make([]string, 0, len(resources))
	for _, r := range resources {
		s := strings.TrimSpace(string(r))
		if s != "" {
			result = append(result, s)
		}
	}
	return result
}


// StringsToResources converts a slice of strings into []Resource
func StringsToResources(strs []string) []Resource {
	res := make([]Resource, len(strs))
	for i, s := range strs {
		res[i] = Resource(s)
	}
	return res
}
