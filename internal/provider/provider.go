package provider

import (
	"context"

	"github.com/abergmeier/terraform-provider-dbussecretservice/internal/login"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	provider := &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"dbus_secretservice_login": login.Resource(),
		},
		Schema: map[string]*schema.Schema{},
	}
	provider.ConfigureContextFunc = func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		// Shameless plug from https://github.com/terraform-providers/terraform-provider-aws/blob/d51784148586f605ab30ecea268e80fe83d415a9/aws/provider.go
		terraformVersion := provider.TerraformVersion
		if terraformVersion == "" {
			// Terraform 0.12 introduced this field to the protocol
			// We can therefore assume that if it's missing it's 0.10 or 0.11
			terraformVersion = "0.11+compatible"
		}
		return nil, nil
	}
	return provider
}
