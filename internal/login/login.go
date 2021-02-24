package login

import (
	"github.com/abergmeier/terraform-provider-dbussecretservice/internal/secretservice"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Resource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"attributes": {
				Type:        schema.TypeMap,
				Required:    true,
				Description: "Attributes to search for",
				ForceNew:    true,
			},
			"parameters": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Algorithm dependent parameters for secret value encoding.",
			},
			"value": {
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
				Description: "Possibly encoded secret value",
			},
			"content_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The content type of the secret. For example: 'text/plain; charset=utf8'",
			},
		},
		Read: read,
	}
}

func read(d *schema.ResourceData, meta interface{}) error {

	tfAttrs := d.Get("attributes").(map[string]interface{})

	attrs := make(map[string]string, len(tfAttrs))
	for k, vi := range tfAttrs {
		v := vi.(string)
		attrs[k] = v
	}

	secrets, err := secretservice.SearchLogin(attrs)
	if err != nil {
		return err
	}

	if len(secrets) == 0 {
		return nil
	}

	d.Set("parameters", secrets[0].Parameters)
	d.Set("content_type", secrets[0].ContentType)
	d.Set("value", secrets[0].Value)
	return nil
}
