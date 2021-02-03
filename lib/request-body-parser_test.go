package lib

import (
	"bytes"
	"fmt"
	"github.com/zclconf/go-cty/cty"
	"io"
	"mime/multipart"
	"os"
	"reflect"
	"strings"
	"testing"
)

func Test_parseBody(t *testing.T) {
	textPlainStr := "hey ho"
	textPlainBody := cty.StringVal(textPlainStr)
	textPlainReader := strings.NewReader(textPlainStr)
	textPlainHeaders := map[string]string{"content-type": "text/plain"}

	applicationJsonStr := `{
  "hey": "ho"
}`
	applicationJsonBody := cty.ObjectVal(map[string]cty.Value{"hey": cty.StringVal("ho")})
	applicationJsonReader := strings.NewReader(applicationJsonStr)
	applicationJsonHeaders := map[string]string{"content-type": "application/json"}

	formUrlencodedStr := "hey=ho"
	formUrlencodedBody := cty.ObjectVal(map[string]cty.Value{"hey": cty.StringVal("ho")})
	formUrlencodedReader := strings.NewReader(formUrlencodedStr)
	formUrlencodedHeaders := map[string]string{"content-type": "application/x-www-form-urlencoded"}

	filename := "/tmp/request-body-parser_test-dummy-1.txt"
	file, _ := os.Create(filename)
	defer os.Remove(filename)
	file.WriteString("test")
	file.Close()

	multipartBody := cty.ObjectVal(map[string]cty.Value{"hey": cty.StringVal("ho"), "file": cty.StringVal(fmt.Sprintf("###READFILE=%s###", filename))})
	multipartBodyHeaders := map[string]string{"content-type": "multipart/form-data"}

	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)
	boundary := "test"
	writer.SetBoundary(boundary)
	fwfile, _ := writer.CreateFormFile("file", "request-body-parser_test-dummy-1.txt")
	fwfile.Write([]byte("test"))

	fwfield, _ := writer.CreateFormField("hey")
	fwfield.Write([]byte("ho"))

	writer.Close()
	multipartBodyReader := bytes.NewReader(buf.Bytes())

	type args struct {
		bodyAsCtyValue cty.Value
		headers        *map[string]string
		boundary       *string
	}

	tests := []struct {
		name    string
		args    args
		want    io.Reader
		wantErr bool
	}{
		{"null body", args{cty.NullVal(cty.String), nil, nil}, nil, false},
		{"text/plain", args{textPlainBody, &textPlainHeaders, nil}, textPlainReader, false},
		{"application/json", args{applicationJsonBody, &applicationJsonHeaders, nil}, applicationJsonReader, false},
		{"application/x-www-form-urlencoded", args{formUrlencodedBody, &formUrlencodedHeaders, nil}, formUrlencodedReader, false},
		{"multipart/form-data", args{multipartBody, &multipartBodyHeaders, &boundary}, multipartBodyReader, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader, err := parseBody(tt.args.bodyAsCtyValue, tt.args.boundary, tt.args.headers)

			if (err != nil) != tt.wantErr {
				t.Errorf("parseBody()  = %v, want %v", err, tt.wantErr)
			} else if tt.want == nil && reader != nil {
				t.Errorf("parseBody() = %v, want %v", reader, tt.want)
			} else if tt.want != nil {
				bufWant := new(bytes.Buffer)
				_, bufWantErr := bufWant.ReadFrom(tt.want)
				if bufWantErr != nil {
					t.Error("parseBody() = want reader error", bufWantErr)
				}

				bufGot := new(bytes.Buffer)
				_, bufGotErr := bufGot.ReadFrom(reader)
				if bufGotErr != nil {
					t.Error("parseBody() = got reader error", bufGotErr)
				}

				if !bytes.Equal(bufWant.Bytes(), bufGot.Bytes()) {
					t.Errorf("parseBody() got;\n%v\nwant;\n%v\n", bufGot.String(), bufWant.String())
				}
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := bodyAsJson(tt.args.bodyAsCtyValue)
			if (err != nil) != tt.wantErr {
				t.Errorf("bodyAsJson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			bufGot := new(bytes.Buffer)
			bufGot.ReadFrom(got)

			bufWant := new(bytes.Buffer)
			bufWant.ReadFrom(tt.want)
			if !reflect.DeepEqual(bufGot.String(), bufWant.String()) {
				t.Errorf("bodyAsJson() got = %v, want %v", bufGot.String(), bufWant.String())
			}
		})
	}
}
