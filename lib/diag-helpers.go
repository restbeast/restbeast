package lib

import (
	"github.com/hashicorp/hcl/v2"
)

func silenceValueMustBeKnown(diags hcl.Diagnostics) hcl.Diagnostics {
	var slicedDiags hcl.Diagnostics

	for _, diag := range diags {
		if !(diag.Summary == "Unsuitable value type" && diag.Detail == "Unsuitable value: value must be known") {
			slicedDiags = slicedDiags.Append(diag)
		}
	}

	return slicedDiags
}
