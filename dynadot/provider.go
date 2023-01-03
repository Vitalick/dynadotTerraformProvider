package dynadot

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vitalick/dynadotTerraformProvider/dynadot/dynoclient"
)

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	domain := d.Get("domain").(string)
	key := d.Get("api_key").(string)

	client := dynoclient.NewDynadotClient(domain, key)

	return client, nil
}

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"dynadot_dns_record": resourceDynadotDNSRecord(),
		},
	}
}
