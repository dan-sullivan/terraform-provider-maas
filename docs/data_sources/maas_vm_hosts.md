
# Data Source: maas_vm_hosts

Provides details about all existing MAAS VM hosts.

## Example Usage

### Using pre-deployed VM host

```terraform
data "maas_vm_hosts" "vm_hosts" {
  id = "rack-01"
}
```

## Argument Reference

The following arguments are supported:

* `id` - (Required) Internal ID. 

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `system_id[]` - The VM host IDs.
* `no[]` - The VM host internal IDs (for create VM Instances)
* `name[]` - The VM host names. 
* `recommended` - The VM host internal ID with the most free memory (for create VM Instances)

