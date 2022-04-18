package maas

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ionutbalutoiu/gomaasclient/client"
)

func dataSourceMaasMachines() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceMachinesRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"system_id": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"hostname": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"zone": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"pool": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"power_type": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"architecture": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"description": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceMachinesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*client.Client)
	machines, err := client.Machines.Get()
	if err != nil {
		return diag.FromErr(err)
	}

	id := d.Get("id").(string)
	system_id := make([]string, len(machines))
	names := make([]string, len(machines))
	zones := make([]string, len(machines))
	pools := make([]string, len(machines))
	power_types := make([]string, len(machines))
	architectures := make([]string, len(machines))
	descriptions := make([]string, len(machines))

	for i, machine := range machines {
		system_id[i] = machine.SystemID
		names[i] = machine.Hostname
		zones[i] = machine.Zone.Name
		pools[i] = machine.Pool.Name
		power_types[i] = machine.PowerType
		architectures[i] = machine.Architecture
		descriptions[i] = machine.Description
	}

	tfState := map[string]interface{}{
		"id":           id,
		"system_id":    system_id,
		"hostname":     names,
		"zone":         zones,
		"pool":         pools,
		"power_type":   power_types,
		"architecture": architectures,
		"description":  descriptions,
	}
	if err := setTerraformState(d, tfState); err != nil {
		return diag.FromErr(err)
	}
	return nil
}
