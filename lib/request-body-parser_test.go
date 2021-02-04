package lib

import (
	"bytes"
	"fmt"
	"github.com/zclconf/go-cty/cty"
	"io"
	"net/url"
	"os"
	"reflect"
	"strings"
	"testing"
)

func Test_parseBody(t *testing.T) {
	textPlainStr := "hey ho"
	textPlainBody := cty.StringVal(textPlainStr)
	textPlainHeaders := Headers{}
	textPlainHeaders.Add("content-type", "text/plain")

	applicationJsonBody := cty.ObjectVal(map[string]cty.Value{"hey": cty.StringVal("ho")})
	applicationJsonHeaders := Headers{}
	applicationJsonHeaders.Add("content-type", "application/json")

	formUrlencodedBody := cty.ObjectVal(map[string]cty.Value{"hey": cty.StringVal("ho")})
	formUrlencodedHeaders := Headers{}
	formUrlencodedHeaders.Add("content-type", "application/x-www-form-urlencoded")

	filename := "/tmp/request-body-parser_test-dummy-1.txt"
	file, _ := os.Create(filename)
	defer os.Remove(filename)
	file.WriteString("test")
	file.Close()

	multipartBody := cty.ObjectVal(
		map[string]cty.Value{
			"hey":   cty.StringVal("ho"),
			"num":   cty.NumberIntVal(int64(10)),
			"boolF": cty.BoolVal(false),
			"boolT": cty.BoolVal(true),
			"file":  cty.StringVal(fmt.Sprintf("###READFILE=%s###", filename)),
		})
	multipartBodyHeaders := Headers{}
	multipartBodyHeaders.Add("content-type", "multipart/form-data; boundary=test")

	filename2 := "/tmp/request-body-parser_test-dummy-1.txt"
	file2, _ := os.Create(filename)
	defer os.Remove(filename)
	file2.WriteString("test")
	file2.Close()
	onlyfileBody := cty.StringVal(fmt.Sprintf("###READFILE=%s###", filename2))
	onlyfileHeaders := Headers{}

	type args struct {
		bodyAsCtyValue cty.Value
		headers        Headers
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"null body", args{cty.NullVal(cty.String), Headers{}}, false},
		{"text/plain", args{textPlainBody, textPlainHeaders}, false},
		{"application/json", args{applicationJsonBody, applicationJsonHeaders}, false},
		{"application/x-www-form-urlencoded", args{formUrlencodedBody, formUrlencodedHeaders}, false},
		{"multipart/form-data", args{multipartBody, multipartBodyHeaders}, false},
		{"file", args{onlyfileBody, onlyfileHeaders}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := parseBody(tt.args.bodyAsCtyValue, tt.args.headers)

			if (err != nil) != tt.wantErr {
				t.Errorf("parseBody()  = %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func Test_bodyAsString(t *testing.T) {
	type args struct {
		bodyAsCtyValue cty.Value
	}
	tests := []struct {
		name    string
		args    args
		want    io.Reader
		wantErr bool
	}{
		{"test", args{cty.StringVal("test")}, strings.NewReader("test"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := bodyAsString(tt.args.bodyAsCtyValue)
			if (err != nil) != tt.wantErr {
				t.Errorf("bodyAsString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bodyAsString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bodyAsNumber(t *testing.T) {
	type args struct {
		bodyAsCtyValue cty.Value
	}
	tests := []struct {
		name    string
		args    args
		want    io.Reader
		wantErr bool
	}{
		{"test", args{cty.NumberIntVal(int64(10))}, strings.NewReader("10"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := bodyAsNumber(tt.args.bodyAsCtyValue)
			if (err != nil) != tt.wantErr {
				t.Errorf("bodyAsNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bodyAsNumber() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bodyAsBool(t *testing.T) {
	type args struct {
		bodyAsCtyValue cty.Value
	}
	tests := []struct {
		name    string
		args    args
		want    io.Reader
		wantErr bool
	}{
		{"test true", args{cty.BoolVal(true)}, strings.NewReader("true"), false},
		{"test false", args{cty.BoolVal(false)}, strings.NewReader("false"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := bodyAsBool(tt.args.bodyAsCtyValue)
			if (err != nil) != tt.wantErr {
				t.Errorf("bodyAsBool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bodyAsBool() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bodyAsJson(t *testing.T) {
	jsonStr := `{
  "key": "value"
}`

	type args struct {
		bodyAsCtyValue cty.Value
	}
	tests := []struct {
		name    string
		args    args
		want    io.Reader
		wantErr bool
	}{
		{"test", args{cty.ObjectVal(map[string]cty.Value{"key": cty.StringVal("value")})}, strings.NewReader(jsonStr), false},
		{"error", args{cty.NegativeInfinity}, strings.NewReader(jsonStr), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := bodyAsJson(tt.args.bodyAsCtyValue)
			if (err != nil) != tt.wantErr {
				t.Errorf("bodyAsJson() error = %v, wantErr %v", err, tt.wantErr)
				return
			} else if !tt.wantErr {
				bufGot := new(bytes.Buffer)
				bufGot.ReadFrom(got)

				bufWant := new(bytes.Buffer)
				bufWant.ReadFrom(tt.want)
				if !reflect.DeepEqual(bufGot.String(), bufWant.String()) {
					t.Errorf("bodyAsJson() got = %v, want %v", bufGot.String(), bufWant.String())
				}
			}
		})
	}
}

func Test_processFormBody(t *testing.T) {
	type args struct {
		parent         *string
		bodyAsCtyValue cty.Value
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    string
	}{
		{"invalid body", args{nil, cty.StringVal("hey-ho")}, true, ""},
		{"has-a-string", args{nil, cty.ObjectVal(map[string]cty.Value{"hey": cty.StringVal("ho")})}, false, "hey=ho"},
		{"has-a-number", args{nil, cty.ObjectVal(map[string]cty.Value{"hey": cty.NumberIntVal(int64(10))})}, false, "hey=10"},
		{"has-a-true-bool", args{nil, cty.ObjectVal(map[string]cty.Value{"hey": cty.BoolVal(true)})}, false, "hey=true"},
		{"has-a-false-bool", args{nil, cty.ObjectVal(map[string]cty.Value{"hey": cty.BoolVal(false)})}, false, "hey=false"},
		{"has-a-object", args{nil, cty.ObjectVal(map[string]cty.Value{"hey": cty.ObjectVal(map[string]cty.Value{"ho": cty.StringVal("no")})})}, false, "hey%5Bho%5D=no"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params := url.Values{}
			err := processFormBody(&params, tt.args.parent, tt.args.bodyAsCtyValue)
			if (err != nil) != tt.wantErr {
				t.Errorf("processFormBody() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_getBoundary(t *testing.T) {
	boundary := "test"
	type args struct {
		contentType string
	}
	tests := []struct {
		name         string
		args         args
		wantBoundary *string
	}{
		{"nil result", args{"multipart/form-data"}, nil},
		{"correct result", args{"multipart/form-data; boundary=test"}, &boundary},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBoundary := getBoundary(tt.args.contentType)
			if (tt.wantBoundary == nil && gotBoundary != nil) || (gotBoundary != nil && !reflect.DeepEqual(*gotBoundary, *tt.wantBoundary)) {
				t.Errorf("getBoundary() = %v, want %v", gotBoundary, tt.wantBoundary)
			}
		})
	}
}

func Test_processFileBody(t *testing.T) {
	type args struct {
		readfileStr string
	}

	filename := "/tmp/request-body-parser_test-dummy-2.txt"
	file, _ := os.Create(filename)
	defer os.Remove(filename)
	file.WriteString("test")
	file.Close()

	tests := []struct {
		name       string
		args       args
		wantReader io.Reader
		wantMime   string
		wantErr    bool
	}{
		{"txt file", args{fmt.Sprintf("###READFILE=%s###", filename)}, strings.NewReader("test"), "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotReader, gotMime, err := processFileBody(tt.args.readfileStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("processFileBody() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var gotBuf bytes.Buffer
			io.Copy(&gotBuf, gotReader)

			var wantBuf bytes.Buffer
			io.Copy(&wantBuf, tt.wantReader)
			if !reflect.DeepEqual(gotBuf.String(), wantBuf.String()) {
				t.Errorf("processFileBody() gotReader = %v, want %v", gotBuf.String(), wantBuf.String())
			}
			if gotMime != tt.wantMime {
				t.Errorf("processFileBody() gotMime = %v, want %v", gotMime, tt.wantMime)
			}
		})
	}
}
