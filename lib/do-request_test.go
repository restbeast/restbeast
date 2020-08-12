package lib

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDoRequest(t *testing.T) {
	set := []struct {
		Request          Request
		ExecutionContext ExecutionContext
	}{
		// TODO: Increase test cases
		{
			Request{
				Method:      "POST",
				Url:         "",
				Headers:     nil,
				Body:        "",
				EvalContext: EvalContext{},
			},
			ExecutionContext{
				Version: "",
				Debug:   true,
			},
		},
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	defer ts.Close()

	for i, run := range set {
		run.Request.Url = fmt.Sprintf("%s/%d", ts.URL, i)
		_, err := DoRequest(run.Request, &run.ExecutionContext)

		if err != nil {
			t.Errorf("do-request failed")
		}
	}
}
