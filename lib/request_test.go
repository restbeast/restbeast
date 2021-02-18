package lib

import (
	"bytes"
	"fmt"
	"github.com/restbeast/restbeast/lib/mocks"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestRequest_Exec(t *testing.T) {
	test1 := mocks.Responder(func(*http.Request) (*http.Response, error) {
		return nil, fmt.Errorf("oh my")
	})
	test2 := mocks.Responder(func(*http.Request) (*http.Response, error) {
		return &http.Response{
			Body:   ioutil.NopCloser(bytes.NewReader([]byte("hello world"))),
			Header: http.Header{"header1": []string{"value1"}},
		}, nil
	})

	var mockTransport = mocks.MockTransport{}
	mockTransport.RegisterResponder("GET", "URL1", test1)
	mockTransport.RegisterResponder("GET", "URL2", test2)

	testHeaders1 := Headers{}
	testHeaders1.Add("header1", "value1")

	type fields struct {
		Method            string
		Url               string
		FullUrl           string
		Headers           Headers
		Body              io.Reader
		Params            *map[string]string
		EvalContext       EvalContext
		PrecedingRequests []*Response
		Response          *Response
		ExecutionContext  *ExecutionContext
		RoundTripper      http.RoundTripper
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "err1",
			fields: fields{
				Method:      "",
				Url:         "URL1",
				FullUrl:     "URL1",
				Headers:     testHeaders1,
				Body:        nil,
				EvalContext: EvalContext{},
				ExecutionContext: &ExecutionContext{
					Version: "v0.0.0-test",
					Debug:   false,
				},
				PrecedingRequests: nil,
				Response:          nil,
				RoundTripper:      &mockTransport,
			},
			wantErr: true,
		},
		{
			name: "test1",
			fields: fields{
				Method:      "GET",
				Url:         "URL1",
				FullUrl:     "URL1",
				Headers:     testHeaders1,
				Body:        nil,
				EvalContext: EvalContext{},
				ExecutionContext: &ExecutionContext{
					Version: "v0.0.0-test",
					Debug:   false,
				},
				PrecedingRequests: nil,
				Response:          nil,
				RoundTripper:      &mockTransport,
			},
			wantErr: true,
		},
		{
			name: "test2",
			fields: fields{
				Method:      "GET",
				Url:         "URL2",
				FullUrl:     "URL2",
				Headers:     testHeaders1,
				Body:        nil,
				EvalContext: EvalContext{},
				ExecutionContext: &ExecutionContext{
					Version: "v0.0.0-test",
					Debug:   true,
				},
				PrecedingRequests: nil,
				Response:          nil,
				RoundTripper:      &mockTransport,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := &Request{
				Method:            tt.fields.Method,
				Url:               tt.fields.Url,
				FullUrl:           tt.fields.FullUrl,
				Headers:           tt.fields.Headers,
				Body:              tt.fields.Body,
				EvalContext:       tt.fields.EvalContext,
				PrecedingRequests: tt.fields.PrecedingRequests,
				Response:          tt.fields.Response,
				ExecutionContext:  tt.fields.ExecutionContext,
				RoundTripper:      tt.fields.RoundTripper,
			}
			if err := request.Exec(); (err != nil) != tt.wantErr {
				t.Errorf("Exec() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRequest_SetUrl(t *testing.T) {
	type args struct {
		urlToSet string
		params   map[string]string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test-1", args{"test-1", map[string]string{}}, "test-1"},
		{"test-2", args{"test-1", map[string]string{"key1": "value1"}}, "test-1?key1=value1"},
		{"test-3", args{"test-1?key1=value1", map[string]string{"key2": "value2"}}, "test-1?key1=value1&key2=value2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := &Request{
				Params: tt.args.params,
			}
			request.SetUrl(tt.args.urlToSet)

			if request.Url != tt.args.urlToSet {
				t.Errorf("SetUrl() got = %s, want %s", request.Url, tt.args.urlToSet)
			}

			if request.FullUrl != tt.want {
				t.Errorf("Exec() got = %s, want %s", request.FullUrl, tt.want)
			}
		})
	}
}

func TestRequest_SetMethod(t *testing.T) {
	type args struct {
		methodToSet string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"post", args{"post"}, "POST"},
		{"get", args{"get"}, "GET"},
		{"empty", args{""}, "GET"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := &Request{}
			request.SetMethod(tt.args.methodToSet)

			if request.Method != tt.want {
				t.Errorf("SetMethod() got = %s, want %s", request.Method, tt.want)
			}
		})
	}
}
