package dynadot

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	//"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var Schema = map[string]*schema.Schema{
	"api_version": {
		Type:     schema.TypeString,
		Optional: true,
		Default:  "",
	},

	"hostname": {
		Type:     schema.TypeString,
		Required: true,
	},

	"headers": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	},
}
