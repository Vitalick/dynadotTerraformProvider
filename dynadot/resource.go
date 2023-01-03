package dynadot

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/vitalick/dynadotTerraformProvider/dynadot/dynorecord"
	//"github.com/terraform-providers/terraform-provider-aws/aws/internal/keyvaluetags"
	//"github.com/terraform-providers/terraform-provider-aws/aws/internal/service/route53/waiter"
)

var Schema = map[string]*schema.Schema{
	"name": {
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
	},

	"type": {
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
		ValidateFunc: validation.StringInSlice([]string{
			"A",
		}, false),
	},

	"zone_id": {
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
	},

	"records": {
		Type:     schema.TypeList,
		Required: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	},

	"ttl": {
		Type:     schema.TypeInt,
		Optional: true,
		Default:  300,
	},

	"set_identifier": {
		Type:     schema.TypeString,
		Optional: true,
	},

	"weight": {
		Type:     schema.TypeInt,
		Optional: true,
	},
	"alias": {
		Type:     schema.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": {
					Type:     schema.TypeString,
					Required: true,
				},
				"zone_id": {
					Type:     schema.TypeString,
					Required: true,
				},
				"evaluate_target_health": {
					Type:     schema.TypeBool,
					Optional: true,
					Default:  false,
				},
			},
		},
	},

	"health_check_id": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"failover": {
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: validation.StringInSlice(failovers, false),
	},

	"multivalue_answer_routing_policy": {
		Type:     schema.TypeBool,
		Optional: true,
	},

	"geolocation_routing_policy": {
		Type:     schema.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"continent": {
					Type:         schema.TypeString,
					Optional:     true,
					ValidateFunc: validation.StringInSlice(continents, false),
				},
				"country": {
					Type:         schema.TypeString,
					Optional:     true,
					ValidateFunc: validation.StringInSlice(countryCodes, false),
				},
				"subdivision": {
					Type:     schema.TypeString,
					Optional: true,
				},
			},
		},
	},
}

func resourceDynadotDNSRecord() *schema.Resource {
	return &schema.Resource{
		CreateContext: dynorecord.ResourceDynadotDNSRecordCreate,
		ReadContext:   dynorecord.ResourceDynadotDNSRecordRead,
		UpdateContext: dynorecord.ResourceDynadotDNSRecordUpdate,
		DeleteContext: dynorecord.ResourceDynadotDNSRecordDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: Schema,
	}
}
