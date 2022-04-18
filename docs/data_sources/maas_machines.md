
# Data Source: maas_machinces

Provides details about all existing MAAS Machines.

## Example Usage

### Get all Machines

```terraform
data "maas_machines" "machines" {
  id  = "rack-01"
}
```

## Argument Reference

The following arguments are supported:

* `id` - (Required) Internal ID. 

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `system_id[]` - The VM host IDs.
* `hostname[]` - The VM host names. 
* `zone[]` - The deployed MAAS machine zone name.
* `pool[]` - The deployed MAAS machine pool name.
* `power_type[]` - (Required) A power management type (e.g. `ipmi`).
* `architecture[]` - The architecture type of the machine. Defaults to `amd64/generic`.
* `description[]` - Description
