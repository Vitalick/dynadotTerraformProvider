package dynorecord

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vitalick/dynadotTerraformProvider/dynadot/dynoclient"
)

func ResourceDynadotDNSRecordDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*dynoclient.DynadotClient)

	name := d.Id()
	zoneID := d.Get("zone_id").(string)

	err := client.DeleteDNS(name, zoneID)
	if err != nil {
		return diag.FromErr(fmt.Errorf("Error deleting Dynadot DNS record: %s", err))
	}

	d.SetId("")
	return nil
}
