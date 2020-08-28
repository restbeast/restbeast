package lib

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
	"net/http"
	"reflect"
	"testing"
)

func Test_lowercaseHeaders(t *testing.T) {
	val1 := []string{
		"header-value-1",
	}
	headers := http.Header{}
	headers["Key1"] = val1

	wantHeaders := headers
	wantHeaders["key1"] = val1

	type args struct {
		headers http.Header
	}

	tests := []struct {
		name string
		args args
		want http.Header
	}{
		{"adds lower case headers", args{headers}, wantHeaders},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowercaseHeaders(tt.args.headers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lowercaseHeaders() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getObjSpec(t *testing.T) {
	tests := []struct {
		name string
		want hcldec.ObjectSpec
	}{
		{"success", hcldec.ObjectSpec{
			"method": &hcldec.AttrSpec{
				Name:     "method",
				Required: true,
				Type:     cty.String,
			},
			"url": &hcldec.AttrSpec{
				Name:     "url",
				Required: true,
				Type:     cty.String,
			},
			"headers": &hcldec.AttrSpec{
				Name:     "headers",
				Required: false,
				Type:     cty.Map(cty.String),
			},
			"body": &hcldec.AttrSpec{
				Name:     "body",
				Required: false,
				Type:     cty.DynamicPseudoType,
			},
			"depends_on": &hcldec.AttrSpec{
				Name:     "depends_on",
				Required: false,
				Type:     cty.List(cty.String),
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getObjSpec(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getObjSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}
