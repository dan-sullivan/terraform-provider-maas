package maas

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ionutbalutoiu/gomaasclient/client"
	"github.com/ionutbalutoiu/gomaasclient/entity"
)

func resourceMaasVMInstance() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceVMInstanceCreate,
		ReadContext:   resourceInstanceRead,
		DeleteContext: resourceVMInstanceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				client := m.(*client.Client)
				machine, err := getMachine(client, d.Id())
				if err != nil {
					return nil, err
				}
				if machine.StatusName != "Deployed" {
					return nil, fmt.Errorf("machine '%s' needs to be already deployed to be imported as maas_instance resource", machine.Hostname)
				}
				d.SetId(machine.SystemID)
				return []*schema.ResourceData{d}, nil
			},
		},

		Schema: map[string]*schema.Schema{
			"kvm_no": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"hostname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},			
			"fqdn": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cpu_count": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
				Default:  1,
			},
			"memory": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
				Default:  2048,
			},
			"storage": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
				Default:  8,
			},
			"zone": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				Default:  nil,
			},
			"pool": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				Default:  nil,
			},
			"tags": {
				Type:     schema.TypeSet,
				Computed: true,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"user_data": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"ip_addresses": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceVMInstanceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*client.Client)

	no := d.Get("kvm_no").(int)
	params := &entity.VMHostMachineParams{
		Cores:    d.Get("cpu_count").(int),
		Memory:   d.Get("memory").(int),
		Storage:  strconv.Itoa(d.Get("storage").(int)),
		Hostname: d.Get("hostname").(string),
	}
	machine, err := client.VMHost.Compose(no, params)
	if err != nil {
		return diag.FromErr(err)
	}

	// Save system id
	d.SetId(machine.SystemID)

	// Wait for MAAS machine to be ready
	_, err = waitForMachineStatus(ctx, client, machine.SystemID, []string{"Commissioning", "Testing"}, []string{"Ready"})
	if err != nil {
		return diag.FromErr(err)
	}

	// Update pool and zone
	mparams := &entity.MachineParams{
		Pool: d.Get("pool").(string),
		Zone: d.Get("zone").(string),
		Description: d.Get("description").(string),
	}
	mpower := make(map[string]string)
	mpower[machine.PowerType] = "virsh"

	_, err = client.Machine.Update(machine.SystemID, mparams, mpower)
	if err != nil {
		return diag.FromErr(err)
	}

	userdata := d.Get("user_data").(string)
	deploy := &entity.MachineDeployParams{
		UserData: userdata,
	}

	// Deploy
	_, err = client.Machine.Deploy(machine.SystemID, deploy)
	if err != nil {
		return diag.FromErr(err)
	}

	// Wait for MAAS machine to be deployed
	_, err = waitForMachineStatus(ctx, client, machine.SystemID, []string{"Deploying"}, []string{"Deployed"})
	if err != nil {
		return diag.FromErr(err)
	}

	// Read MAAS machine info
	return resourceInstanceRead(ctx, d, m)

}

func resourceVMInstanceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*client.Client)

	// Delete VM MAAS machine
	err := client.Machine.Delete(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
