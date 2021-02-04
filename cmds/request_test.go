package cmds

import (
	"fmt"
	"github.com/restbeast/restbeast/lib"
	"net/http"
	"testing"
	"time"
)

func Test_printJustTiming(t *testing.T) {
	type args struct {
		response lib.Response
		padding  string
	}

	resp1 := lib.Response{
		Method:        "POST",
		Url:           "http://localhost",
		StatusCode:    0,
		Proto:         "",
		Body:          nil,
		Headers:       nil,
		Timing:        lib.RequestTiming{},
		Request:       nil,
		BytesSend:     0,
		BytesReceived: 10,
	}
	matchString1 := `POST http://localhost
  │  Total Time: 0 ms
  │  Bytes Received: 10 B
`

	resp2 := lib.Response{
		Method:        "POST",
		Url:           "http://localhost",
		StatusCode:    0,
		Proto:         "",
		Body:          nil,
		Headers:       nil,
		Timing:        lib.RequestTiming{},
		Request:       nil,
		BytesSend:     10,
		BytesReceived: 10,
	}
	matchString2 := `POST http://localhost
  │  Total Time: 0 ms
  │  Bytes Sent: 10 B
  │  Bytes Received: 10 B
`

	resp3 := lib.Response{
		Method:        "POST",
		Url:           "http://localhost",
		StatusCode:    0,
		Proto:         "",
		Body:          nil,
		Headers:       nil,
		Timing:        lib.RequestTiming{},
		Request:       nil,
		BytesSend:     10,
		BytesReceived: 10,
	}
	matchString3 := `POST http://localhost
    │  Total Time: 0 ms
    │  Bytes Sent: 10 B
    │  Bytes Received: 10 B
`

	tests := []struct {
		name string
		args args
		want string
	}{
		{"get requests", args{resp1, ""}, matchString1},
		{"requests with payload", args{resp2, ""}, matchString2},
		{"with padding", args{resp3, "  "}, matchString3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			returnVal := printJustTiming(tt.args.response, tt.args.padding)

			if returnVal != tt.want {
				t.Errorf("printJustTiming() got1 = %v, want %v", returnVal, tt.want)
			}
		})
	}
}

func Test_printDetailedTiming(t *testing.T) {
	type args struct {
		response lib.Response
		padding  string
	}

	resp1 := lib.Response{
		Method:        "POST",
		Url:           "http://localhost",
		StatusCode:    0,
		Proto:         "",
		Body:          nil,
		Headers:       nil,
		Timing:        lib.RequestTiming{},
		Request:       nil,
		BytesSend:     0,
		BytesReceived: 10,
	}
	matchString1 := `POST http://localhost
  │  DNS Resolve Time: 0 ms
  │  Connection Time: 0 ms
  │  First Byte Time: 0 ms
  │  Total Time: 0 ms
  │  Bytes Received: 10 B
`

	resp2 := lib.Response{
		Method:        "POST",
		Url:           "http://localhost",
		StatusCode:    0,
		Proto:         "",
		Body:          nil,
		Headers:       nil,
		Timing:        lib.RequestTiming{},
		Request:       nil,
		BytesSend:     10,
		BytesReceived: 10,
	}
	matchString2 := `POST http://localhost
  │  DNS Resolve Time: 0 ms
  │  Connection Time: 0 ms
  │  First Byte Time: 0 ms
  │  Total Time: 0 ms
  │  Bytes Sent: 10 B
  │  Bytes Received: 10 B
`

	resp3 := lib.Response{
		Method:     "POST",
		Url:        "http://localhost",
		StatusCode: 0,
		Proto:      "",
		Body:       nil,
		Headers:    nil,
		Timing: lib.RequestTiming{
			Dns:       10,
			Conn:      10,
			Tls:       0,
			FirstByte: 10,
			Total:     10,
		},
		Request:       nil,
		BytesSend:     10,
		BytesReceived: 10,
	}
	matchString3 := `POST http://localhost
  │  DNS Resolve Time: 0 ms
  │  Connection Time: 0 ms
  │  First Byte Time: 0 ms
  │  Total Time: 0 ms
  │  Bytes Sent: 10 B
  │  Bytes Received: 10 B
`

	resp4 := lib.Response{
		Method:     "POST",
		Url:        "http://localhost",
		StatusCode: 0,
		Proto:      "",
		Body:       nil,
		Headers:    nil,
		Timing: lib.RequestTiming{
			Dns:       10,
			Conn:      10,
			Tls:       10,
			FirstByte: 10,
			Total:     10,
		},
		Request:       nil,
		BytesSend:     10,
		BytesReceived: 10,
	}
	matchString4 := `POST http://localhost
  │  DNS Resolve Time: 0 ms
  │  Connection Time: 0 ms
  │  TLS Handshake Time: 0 ms
  │  First Byte Time: 0 ms
  │  Total Time: 0 ms
  │  Bytes Sent: 10 B
  │  Bytes Received: 10 B
`

	resp5 := lib.Response{
		Method:     "POST",
		Url:        "http://localhost",
		StatusCode: 0,
		Proto:      "",
		Body:       nil,
		Headers:    nil,
		Timing: lib.RequestTiming{
			Dns:       10,
			Conn:      10,
			Tls:       10,
			FirstByte: 10,
			Total:     10,
		},
		Request:       nil,
		BytesSend:     10,
		BytesReceived: 10,
	}
	matchString5 := `POST http://localhost
    │  DNS Resolve Time: 0 ms
    │  Connection Time: 0 ms
    │  TLS Handshake Time: 0 ms
    │  First Byte Time: 0 ms
    │  Total Time: 0 ms
    │  Bytes Sent: 10 B
    │  Bytes Received: 10 B
`

	tests := []struct {
		name string
		args args
		want string
	}{
		{"get requests", args{resp1, ""}, matchString1},
		{"requests with payload", args{resp2, ""}, matchString2},
		{"with timing", args{resp3, ""}, matchString3},
		{"with tls timing", args{resp4, ""}, matchString4},
		{"with padding", args{resp5, "  "}, matchString5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			returnVal := printDetailedTiming(tt.args.response, tt.args.padding)

			if returnVal != tt.want {
				t.Errorf("printDetailedTiming() got1 = %v, want %v", returnVal, tt.want)
			}
		})
	}
}

func Test_printTiming(t *testing.T) {
	type args struct {
		outputTiming         bool
		outputDetailedTiming bool
		request              lib.Request
		response             lib.Response
		padding              string
	}

	req1 := lib.Request{
		PrecedingRequests: nil,
	}
	resp1 := lib.Response{
		Timing:        lib.RequestTiming{},
		BytesSend:     0,
		BytesReceived: 10,
	}
	matchString1 := ``

	req2 := lib.Request{
		PrecedingRequests: nil,
	}
	resp2 := lib.Response{
		Timing:        lib.RequestTiming{},
		BytesSend:     0,
		BytesReceived: 10,
	}
	matchString2 := `  │
  ├── 
  │  Total Time: 0 ms
  │  Bytes Received: 10 B
`

	req3 := lib.Request{
		PrecedingRequests: nil,
	}
	resp3 := lib.Response{
		Timing:        lib.RequestTiming{},
		BytesSend:     0,
		BytesReceived: 10,
	}
	matchString3 := `  │
  ├── 
  │  DNS Resolve Time: 0 ms
  │  Connection Time: 0 ms
  │  First Byte Time: 0 ms
  │  Total Time: 0 ms
  │  Bytes Received: 10 B
`

	precedingRes4 := &lib.Response{
		Request:       &lib.Request{},
		Timing:        lib.RequestTiming{},
		BytesSend:     0,
		BytesReceived: 10,
	}
	req4 := lib.Request{
		PrecedingRequests: []*lib.Response{precedingRes4},
	}
	resp4 := lib.Response{
		Timing:        lib.RequestTiming{},
		BytesSend:     0,
		BytesReceived: 10,
	}
	matchString4 := `  │
  ├── 
  │  Total Time: 0 ms
  │  Bytes Received: 10 B
  │    │
  │    ├── 
  │    │  Total Time: 0 ms
  │    │  Bytes Received: 10 B
`

	tests := []struct {
		name string
		args args
		want string
	}{
		{"empty output", args{false, false, req1, resp1, ""}, matchString1},
		{"just timing", args{true, false, req2, resp2, ""}, matchString2},
		{"with detailed-timing", args{false, true, req3, resp3, ""}, matchString3},
		{"just timing with preceding", args{true, false, req4, resp4, ""}, matchString4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			returnVal := printTiming(tt.args.outputTiming, tt.args.outputDetailedTiming, tt.args.request, tt.args.response, tt.args.padding)

			if returnVal != tt.want {
				t.Errorf("printTiming() got1 = %v, want %v", returnVal, tt.want)
			}
		})
	}
}

func Test_printHeaders(t *testing.T) {
	headers1 := http.Header{}
	headers1.Set("header-key", "header-value")
	time.Sleep(100 * time.Millisecond)
	headers1.Set("an-other-header-key", "an-other-header-value")

	match := fmt.Sprintf("\n\u001B[1m%s\u001B[0m: %s\n\u001B[1m%s\u001B[0m: %s\n", "Header-Key", "header-value", "An-Other-Header-Key", "an-other-header-value")

	tests := []struct {
		name    string
		headers http.Header
		show    bool
		want    string
	}{
		{"show headers", headers1, true, match},
		{"show none", headers1, false, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			showHeaders = tt.show
			returnVal := printHeaders(lib.Response{Headers: headers1})

			if returnVal != tt.want {
				t.Errorf("printHeaders() got1 = %s, want %s", returnVal, tt.want)
			}
		})
	}
}
