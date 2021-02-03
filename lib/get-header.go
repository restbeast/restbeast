package lib

import (
	"strings"
)

func getHeader(searchKey string, headers *map[string]string) (contentType *string) {
	if headers == nil {
		return nil
	}

	for key, val := range *headers {
		if strings.ToLower(searchKey) == strings.ToLower(key) {
			contentType = &val
		}
	}

	return contentType
}

func getHeaderKey(searchKey string, headers *map[string]string) *string {
	if headers == nil {
		return nil
	}

	for key, _ := range *headers {
		if strings.ToLower(searchKey) == strings.ToLower(key) {
			return &key
		}
	}

	return nil
}
