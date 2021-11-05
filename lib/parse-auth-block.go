package lib

import (
	b64 "encoding/base64"
	. "fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

func parseBasicAuth(request *Request, basicAuth BasicAuthCfg, ctx hcl.EvalContext) hcl.Diagnostics {
	cfg, diags := hcldec.Decode(
		basicAuth.Body, &hcldec.ObjectSpec{
			"username": &hcldec.AttrSpec{
				Name:     "username",
				Required: true,
				Type:     cty.String,
			},
			"password": &hcldec.AttrSpec{
				Name:     "password",
				Required: true,
				Type:     cty.String,
			},
		}, &ctx,
	)

	if diags != nil {
		return diags
	}

	username := cfg.GetAttr("username").AsString()
	password := cfg.GetAttr("password").AsString()

	authString := Sprintf("%s:%s", username, password)
	encodedString := b64.StdEncoding.EncodeToString([]byte(authString))

	request.Headers.Add("Authorization", Sprintf("Basic %s", encodedString))

	return nil
}

func parseBearerAuth(request *Request, basicAuth BearerAuthCfg, ctx hcl.EvalContext) hcl.Diagnostics {
	cfg, diags := hcldec.Decode(
		basicAuth.Body, &hcldec.ObjectSpec{
			"token": &hcldec.AttrSpec{
				Name:     "token",
				Required: true,
				Type:     cty.String,
			},
		}, &ctx,
	)

	if diags != nil {
		return diags
	}

	token := cfg.GetAttr("token").AsString()
	request.Headers.Add("Authorization", Sprintf("Bearer %s", token))

	return nil
}

func parseAuthBlock(request *Request, authBlock *AuthCfg, ctx hcl.EvalContext) hcl.Diagnostics {
	if authBlock != nil {
		if authBlock.BasicAuth != nil {
			diags := parseBasicAuth(request, *authBlock.BasicAuth, ctx)

			if diags != nil {
				return diags
			}
		} else if authBlock.BearerAuth != nil {
			diags := parseBearerAuth(request, *authBlock.BearerAuth, ctx)

			if diags != nil {
				return diags
			}
		}
	}

	return nil
}
