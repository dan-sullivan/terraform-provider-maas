package maas

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ionutbalutoiu/gomaasclient/client"
)

func dataSourceMaasVMHosts() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVMHostsRead,

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
			"no": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"name": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"recommended": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceVMHostsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*client.Client)
	vmHosts, err := client.VMHosts.Get()
	if err != nil {
		return diag.FromErr(err)
	}

	id := d.Get("id").(string)
	numbers := make([]int, len(vmHosts))
	system_id := make([]string, len(vmHosts))
	names := make([]string, len(vmHosts))

	recommended := 0
	memory := 0
	for i, vmHost := range vmHosts {
		numbers[i] = vmHost.ID
		system_id[i] = vmHost.Host.SystemID
		names[i] = vmHost.Name
		if vmHost.Available.Memory > memory {
			recommended = vmHost.ID
			memory = vmHost.Available.Memory
		}
	}

	tfState := map[string]interface{}{
		"id":          id,
		"system_id":   system_id,
		"no":          numbers,
		"name":        names,
		"recommended": recommended,
	}
	if err := setTerraformState(d, tfState); err != nil {
		return diag.FromErr(err)
	}
	return nil

}
