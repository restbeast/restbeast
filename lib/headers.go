package lib

import (
	"fmt"
	"net/http"
	"regexp"
	"sort"
	"strings"
)

type Headers struct {
	kv map[string][]string
}

type OrderedMapCallback func(k, v string)

func (headers *Headers) AddBulk(all map[string]string) *Headers {
	if headers.kv == nil {
		headers.kv = make(map[string][]string)
	}

	for key, val := range all {
		if headers.kv[key] == nil {
			headers.kv[key] = []string{val}
		} else {
			headers.kv[key] = append(headers.kv[key], val)
		}
	}

	return headers
}

func (headers *Headers) Add(k string, v string) *Headers {
	if headers.kv == nil {
		headers.kv = make(map[string][]string)
	}

	if headers.kv[k] == nil {
		headers.kv[k] = []string{v}
	} else {
		headers.kv[k] = append(headers.kv[k], v)
	}

	return headers
}

func (headers *Headers) Set(k string, v string) *Headers {
	if headers.kv == nil {
		headers.kv = make(map[string][]string)
	}

	existingKey := headers.GetKey(k)
	if existingKey != nil {
		headers.kv[*existingKey][0] = v
	} else {
		headers.kv[k] = []string{v}
	}

	return headers
}

func (headers *Headers) Get(searchKey string) (found *string) {
	if headers.kv == nil {
		return nil
	}

	for key, val := range headers.kv {
		if strings.ToLower(searchKey) == strings.ToLower(key) {
			return &val[0]
		}
	}

	return nil
}

func (headers *Headers) GetKey(searchKey string) *string {
	if headers.kv == nil {
		return nil
	}

	for key := range headers.kv {
		if strings.ToLower(searchKey) == strings.ToLower(key) {
			return &key
		}
	}

	return nil
}

func (headers *Headers) GetAll() map[string][]string {
	allHeaders := make(map[string][]string)

	for key, val := range headers.kv {
		allHeaders[key] = val
		allHeaders[strings.ToLower(key)] = val
	}

	return allHeaders
}

func (headers *Headers) SetCookies(cookies map[string]string) {
	cookieRegex := regexp.MustCompile(`;\s?`)

	existingCookieHeader := headers.Get("cookie")
	if existingCookieHeader != nil {
		cookiesFromHeader := cookieRegex.Split(*existingCookieHeader, -1)
		for _, cookie := range cookiesFromHeader {
			kv := strings.Split(cookie, "=")
			if len(kv) == 2 {
				cookies[kv[0]] = kv[1]
			}
		}
	}

	var cookieHeader string
	for k, v := range cookies {
		cookieHeader += fmt.Sprintf("%s=%s; ", k, v)
	}
	headers.Set("cookie", strings.TrimSuffix(cookieHeader, "; "))
}

func (headers *Headers) ToRequest(httpReq *http.Request) {
	if headers.kv == nil {
		return
	}

	for key, value := range headers.kv {
		httpReq.Header.Set(key, value[0])
	}
}

func (headers *Headers) FromResponse(resHeaders http.Header) *Headers {
	if headers.kv == nil {
		headers.kv = make(map[string][]string)
	}

	for key, slice := range resHeaders {
		for i, val := range slice {
			if headers.kv[key] == nil {
				headers.kv[key] = make([]string, len(slice))
			}
			headers.kv[key][i] = val
		}
	}

	return headers
}

func (headers *Headers) OrderedCallBack(cb OrderedMapCallback) {
	if headers.kv == nil {
		headers.kv = make(map[string][]string)
	}

	keys := make([]string, 0, len(headers.kv))
	for k := range headers.kv {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		cb(k, headers.kv[k][0])
	}
}
