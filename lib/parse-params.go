package lib

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
)

func parseParamsBlock(request *Request, paramsBlock *ParamsBlockCfg, ctx hcl.EvalContext) hcl.Diagnostics {
	if paramsBlock != nil {
		var params map[string]string
		diags := gohcl.DecodeBody(paramsBlock.Body, &ctx, &params)
		diags = silenceValueMustBeKnown(diags)

		if diags != nil {
			return diags
		}

		request.Params = params

		return nil
	}

	return nil
}
