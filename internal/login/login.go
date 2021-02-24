package login

import (
	"strings"

	"github.com/abergmeier/terraform-provider-dbussecretservice/internal/datasource"
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

	attrStrings := make([]string, 0, len(tfAttrs))
	attrs := make(map[string]string, len(tfAttrs))
	for k, vi := range tfAttrs {
		v := vi.(string)
		attrs[k] = v
		attrStrings = append(attrStrings, k+":"+v)
	}

	secrets, err := secretservice.SearchLogin(attrs)
	if err != nil {
		return err
	}

	if len(secrets) == 0 {
		return nil
	}

	if secrets[0].Parameters == nil {
		err = d.Set("parameters", nil)
	} else {
		err = d.Set("parameters", string(secrets[0].Parameters))
	}

	if err != nil {
		return err
	}

	err = d.Set("content_type", secrets[0].ContentType)
	if err != nil {
		return err
	}

	if secrets[0].Value == nil {
		err = d.Set("value", nil)
	} else {
		err = d.Set("value", string(secrets[0].Value))
	}

	combinedAttrs := strings.Join(append(attrStrings, datasource.AlwaysUniqueID()), "_")
	d.SetId(combinedAttrs)
	return nil
}
