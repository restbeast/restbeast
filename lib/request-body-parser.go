package lib

import (
	"bytes"
	"encoding/json"
	. "fmt"
	"github.com/h2non/filetype"
	"github.com/zclconf/go-cty/cty"
	ctyjson "github.com/zclconf/go-cty/cty/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/url"
	"path/filepath"
	"regexp"
	"strings"
)

func bodyAsString(bodyAsCtyValue cty.Value) (io.Reader, error) {
	return strings.NewReader(bodyAsCtyValue.AsString()), nil
}

func bodyAsNumber(bodyAsCtyValue cty.Value) (io.Reader, error) {
	return strings.NewReader(bodyAsCtyValue.AsBigFloat().String()), nil
}

func bodyAsBool(bodyAsCtyValue cty.Value) (io.Reader, error) {
	var strVal string
	if bodyAsCtyValue.True() {
		strVal = "true"
	} else {
		strVal = "false"
	}
	return strings.NewReader(strVal), nil
}

func bodyAsJson(bodyAsCtyValue cty.Value) (io.Reader, error) {
	bodyJSON, jsonErr := json.MarshalIndent(ctyjson.SimpleJSONValue{bodyAsCtyValue}, "", "  ")
	if jsonErr != nil {
		return nil, Errorf("Error: failed to parse request body, \n%s\n", jsonErr)
	}

	return bytes.NewReader(bodyJSON), nil
}

func getBoundary(contentType string) (boundary *string) {
	boundaryRegex := regexp.MustCompile(";\\s?boundary=([\\w\\d\\-]+)?$")
	boundaryMatches := boundaryRegex.FindStringSubmatch(contentType)
	if len(boundaryMatches) > 0 {
		boundary = &boundaryMatches[1]
	}

	return boundary
}

func processFormBody(params *url.Values, parent *string, bodyAsCtyValue cty.Value) error {
	bodyType := bodyAsCtyValue.Type()

	if !bodyType.IsObjectType() {
		return Errorf("request body has to be a key/value pairs to use application/x-www-form-urlencoded")
	}

	for k, v := range bodyAsCtyValue.AsValueMap() {
		if parent != nil {
			k = Sprintf("%s[%s]", *parent, k)
		}

		switch v.Type().FriendlyName() {
		case "number":
			params.Add(k, v.AsBigFloat().String())
		case "string":
			params.Add(k, v.AsString())
		case "bool":
			var strVal string
			if v.True() {
				strVal = "true"
			} else {
				strVal = "false"
			}
			params.Add(k, strVal)
		case "object":
			err := processFormBody(params, &k, v)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func processMultipartFormBody(bodyAsCtyValue cty.Value, boundary *string) (io.Reader, string, error) {
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	if boundary != nil {
		err := writer.SetBoundary(*boundary)
		if err != nil {
			return nil, "", err
		}
	}

	for k, v := range bodyAsCtyValue.AsValueMap() {
		err := processMultipartBodyPart(k, v, writer)
		if err != nil {
			return nil, "", err
		}
	}

	err := writer.Close()
	if err != nil {
		return nil, "", err
	}

	return bytes.NewReader(buf.Bytes()), writer.FormDataContentType(), nil
}

// It's impossible to make this function concurrent
// According to https://golang.org/pkg/mime/multipart/#Writer.CreatePart
// > After calling CreatePart, any previous part may no longer be written to.
func processMultipartBodyPart(k string, v cty.Value, writer *multipart.Writer) error {
	switch v.Type().FriendlyName() {
	case "number":
		fw, _ := writer.CreateFormField(k)
		_, _ = io.Copy(fw, strings.NewReader(v.AsBigFloat().String()))

	case "string":
		vStr := v.AsString()

		if strings.HasPrefix(vStr, "###READFILE=") && strings.HasSuffix(vStr, "###") {
			trimmed := strings.TrimPrefix(vStr, "###READFILE=")
			path := strings.TrimSuffix(trimmed, "###")

			file := filepath.Base(path)
			fw, _ := writer.CreateFormFile(k, file)
			contents, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			_, err = fw.Write(contents)
			if err != nil {
				return err
			}
		} else {
			fw, _ := writer.CreateFormField(k)
			_, _ = io.Copy(fw, strings.NewReader(vStr))
		}

	case "bool":
		var strVal string
		if v.True() {
			strVal = "true"
		} else {
			strVal = "false"
		}

		fw, _ := writer.CreateFormField(k)
		_, _ = io.Copy(fw, strings.NewReader(strVal))
	}

	return nil
}

func processFileBody(readfileStr string) (reader io.Reader, mime string, err error) {
	trimmed := strings.TrimPrefix(readfileStr, "###READFILE=")
	path := strings.TrimSuffix(trimmed, "###")

	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, "", err
	}

	match, err := filetype.Match(contents)
	if err != nil {
		return nil, "", err
	} else if match != filetype.Unknown {
		mime = match.MIME.Value
	}

	return bytes.NewReader(contents), mime, nil
}

func parseBody(bodyAsCtyValue cty.Value, headers Headers) (io.Reader, error) {
	contentTypeHeader := headers.Get("content-type")
	var contentType string
	if contentTypeHeader != nil {
		contentType = strings.ToLower(*contentTypeHeader)
	}

	isJson := strings.HasPrefix(contentType, "application/json")
	isForm := strings.HasPrefix(contentType, "application/x-www-form-urlencoded")
	isMultipart := strings.HasPrefix(contentType, "multipart/form-data") ||
		strings.HasPrefix(contentType, "multipart/mixed")
	isFile := bodyAsCtyValue.Type().FriendlyName() == "string" &&
		!bodyAsCtyValue.IsNull() &&
		strings.HasPrefix(bodyAsCtyValue.AsString(), "###READFILE=") &&
		strings.HasSuffix(bodyAsCtyValue.AsString(), "###")

	if !bodyAsCtyValue.IsNull() {
		switch {
		case isJson:
			return bodyAsJson(bodyAsCtyValue)

		case isForm:
			params := url.Values{}
			err := processFormBody(&params, nil, bodyAsCtyValue)
			if err != nil {
				return nil, err
			}
			return strings.NewReader(params.Encode()), nil

		case isMultipart:
			bodyType := bodyAsCtyValue.Type()
			if !bodyType.IsObjectType() {
				return nil, Errorf("request body has to be a key/value pairs to use multipart/form-data")
			}

			contentTypeHeaderKey := *headers.GetKey("content-type")
			boundary := getBoundary(contentType)
			reader, newHeader, err := processMultipartFormBody(bodyAsCtyValue, boundary)
			if err != nil {
				return nil, err
			}

			headers.Add(contentTypeHeaderKey, newHeader)
			return reader, nil

		case isFile:
			reader, mime, err := processFileBody(bodyAsCtyValue.AsString())
			if err != nil {
				return nil, err
			}

			if contentType == "" {
				headers.Add("content-type", mime)
			}

			return reader, nil

		// unknown header will be treated as
		// text/plain if it's number, bool or string
		// application/json if it isn't primitive
		default:
			bodyType := bodyAsCtyValue.Type()

			if bodyType.IsPrimitiveType() {
				switch bodyType.FriendlyName() {
				case "number":
					return bodyAsNumber(bodyAsCtyValue)
				case "string":
					return bodyAsString(bodyAsCtyValue)
				case "bool":
					return bodyAsBool(bodyAsCtyValue)
				}
			} else {
				return bodyAsJson(bodyAsCtyValue)
			}
		}
	}

	return nil, nil
}
