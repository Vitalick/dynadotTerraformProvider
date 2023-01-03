package dynorecord

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vitalick/dynadotTerraformProvider/dynadot/dynoclient"
)

func ResourceDynadotDNSRecordCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*dynoclient.DynadotClient)

	record := &DynadotRecord{
		Name:     d.Get("name").(string),
		Type:     d.Get("type").(string),
		ZoneID:   d.Get("zone_id").(string),
		Records:  expandStringList(d.Get("records").([]interface{})),
		TTL:      d.Get("ttl").(int),
		SetIdent: d.Get("set_identifier").(string),
		Weight:   d.Get("weight").(int),
		Failover: d.Get("failover").(string),
		HealthID: d.Get("health_check_id").(string),
	}

	if v, ok := d.GetOk("multivalue_answer_routing_policy"); ok {
		record.MultiValueAnswer = v.(bool)
	}

	if v, ok := d.GetOk("alias"); ok {
		alias := v.([]interface{})[0].(map[string]interface{})
		record.Alias = &DynadotRecord{
			Name:             alias["name"].(string),
			ZoneID:           alias["zone_id"].(string),
			EvaluateTarget:   alias["evaluate_target_health"].(bool),
			MultiValueAnswer: true,
		}
	}

	if v, ok := d.GetOk("geolocation_routing_policy"); ok {
		geo := v.([]interface{})[0].(map[string]interface{})
		record.GeoLocation = &DynadotGeoLocation{
			Continent:   geo["continent"].(string),
			Country:     geo["country"].(string),
			Subdivision: geo["subdivision"].(string),
		}
	}

	_, err := client.CreateDNS(record)
	if err != nil {
		return diag.FromErr(fmt.Errorf("Error creating Dynadot DNS record: %s", err))
	}

	d.SetId(record.Name)

	return ResourceDynadotDNSRecordRead(ctx, d, meta)
}
