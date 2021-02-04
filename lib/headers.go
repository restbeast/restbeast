package lib

import (
	"net/http"
	"strings"
)

type Headers struct {
	kv map[string]string
}

func (headers *Headers) AddBulk(all map[string]string) {
	if headers.kv == nil {
		headers.kv = make(map[string]string)
	}

	for key, val := range all {
		headers.kv[key] = val
	}
}

func (headers *Headers) Add(k string, v string) *Headers {
	if headers.kv == nil {
		headers.kv = make(map[string]string)
	}

	headers.kv[k] = v

	return headers
}

func (headers *Headers) Get(searchKey string) (found *string) {
	if headers.kv == nil {
		return nil
	}

	for key, val := range headers.kv {
		if strings.ToLower(searchKey) == strings.ToLower(key) {
			found = &val
		}
	}

	return found
}

func (headers *Headers) GetKey(searchKey string) (key *string) {
	if headers.kv == nil {
		return nil
	}

	for key, _ := range headers.kv {
		if strings.ToLower(searchKey) == strings.ToLower(key) {
			return &key
		}
	}

	return nil
}

func (headers *Headers) ToRequest(httpReq *http.Request) {
	if headers.kv == nil {
		return
	}

	for key, value := range headers.kv {
		httpReq.Header.Set(key, value)
	}
}
