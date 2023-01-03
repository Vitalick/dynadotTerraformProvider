package dynorecord

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vitalick/dynadotTerraformProvider/dynadot/dynoclient"
)

func ResourceDynadotDNSRecordRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*dynoclient.DynadotClient)

	name := d.Id()
	zoneID := d.Get("zone_id").(string)

	record, err := client.GetDNS(name, zoneID)
	if err != nil {
		return diag.FromErr(fmt.Errorf("Error retrieving Dynadot DNS record: %s", err))
	}

	d.Set("name", record.Name)
	d.Set("type", record.Type)
	d.Set("zone_id", record.ZoneID)
	d.Set("records", record.Records)
	d.Set("ttl", record.TTL)
	d.Set("set_identifier", record.SetIdent)
	d.Set("weight", record.Weight)
	d.Set("failover", record.Failover)
	d.Set("health_check_id", record.HealthID)
	d.Set("multivalue_answer_routing_policy", record.MultiValueAnswer)

	if record.GeoLocation != nil {
		d.Set("geolocation_routing_policy", []map[string]interface{}{
			{
				"continent":   record.GeoLocation.Continent,
				"country":     record.GeoLocation.Country,
				"subdivision": record.GeoLocation.Subdivision,
			},
		})
	}
	if record.Alias != nil {
		d.Set("alias", []map[string]interface{}{
			{
				"name":                      record.Alias.Name,
				"zone_id":                   record.Alias.ZoneID,
				"evaluate_target_health":    record.Alias.EvaluateTarget,
				"multivalue_answer_routing": record.Alias.MultiValueAnswer,
			},
		})
	}

	return nil
}
