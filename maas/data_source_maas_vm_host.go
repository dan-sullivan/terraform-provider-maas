package maas

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ionutbalutoiu/gomaasclient/client"
)

func dataSourceMaasVMHost() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVMHostRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"no": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cores": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"memory": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"local_storage": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceVMHostRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	vmHost, err := getVMHost(m.(*client.Client), d.Get("name").(string))
	if err != nil {
		return diag.FromErr(err)
	}
	tfState := map[string]interface{}{
		"id":            vmHost.Host.SystemID,
		"no":            vmHost.ID,
		"name":          vmHost.Name,
		"cores":         vmHost.Available.Cores,
		"memory":        vmHost.Available.Memory,
		"local_storage": vmHost.Available.LocalStorage,
	}
	if err := setTerraformState(d, tfState); err != nil {
		return diag.FromErr(err)
	}
	return nil

}
