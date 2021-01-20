package lib

import (
	"github.com/hashicorp/hcl/v2"
	"reflect"
	"testing"
)

func Test_silenceValueMustBeKnown(t *testing.T) {
	diags1 := hcl.Diagnostics{
		&hcl.Diagnostic{
			Severity: 0,
			Summary:  "Nothing to see here",
			Detail:   "Go on",
		},
	}

	diags2 := hcl.Diagnostics{
		&hcl.Diagnostic{
			Severity: 0,
			Summary:  "Unsuitable value type",
			Detail:   "Unsuitable value: value must be known",
		},
	}

	tests := []struct {
		name  string
		diags hcl.Diagnostics
		want  hcl.Diagnostics
	}{
		{
			"nothing to silence",
			diags1,
			diags1,
		},
		{
			"something to silence",
			diags2,
			hcl.Diagnostics{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := silenceValueMustBeKnown(tt.diags)

			if !reflect.DeepEqual(got.Error(), tt.want.Error()) {
				t.Errorf("silenceValueMustBeKnown() = %v, want %v", got, tt.want)
			}
		})
	}
}
