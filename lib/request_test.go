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

	type fields struct {
		Method            string
		Url               string
		Headers           map[string]string
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
				Headers:     map[string]string{"header1": "value1"},
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
				Headers:     map[string]string{"header1": "value1"},
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
				Headers:     map[string]string{"header1": "value1"},
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
