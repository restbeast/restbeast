package lib

import (
	"bytes"
	"github.com/zclconf/go-cty/cty"
	"io"
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
	applicationJsonBody := cty.MapVal(map[string]cty.Value{"hey": cty.StringVal("ho")})
	applicationJsonReader := strings.NewReader(applicationJsonStr)
	applicationJsonHeaders := map[string]string{"content-type": "application/json"}

	type args struct {
		bodyAsCtyValue cty.Value
		headers        *map[string]string
	}

	tests := []struct {
		name    string
		args    args
		want    io.Reader
		wantErr bool
	}{
		{"null body", args{cty.NullVal(cty.String), nil}, nil, false},
		{"text/plain", args{textPlainBody, &textPlainHeaders}, textPlainReader, false},
		{"application/json", args{applicationJsonBody, &applicationJsonHeaders}, applicationJsonReader, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader, err := parseBody(tt.args.bodyAsCtyValue, tt.args.headers)

			if (err != nil) != tt.wantErr {
				t.Errorf("parseBody() = %v, want %v", err, tt.wantErr)
			}

			if tt.want == nil && reader != nil {
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
					t.Errorf("parseBody() = %v, want %v", reader, tt.want)
				}
			}
		})
	}
}
