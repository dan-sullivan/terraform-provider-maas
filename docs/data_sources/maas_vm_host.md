
# Data Source: maas_vm_host

Provides details about an existing MAAS VM hosts.

## Example Usage

### Using pre-deployed VM host

```terraform
data "maas_vm_host" "kvm-01" {
  name = "kvm-01"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The VM host name. 

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The VM host ID.
* `no` - The VM host internal ID (for create VM Instances)
* `cores` - The VM host available number of CPU cores.
* `memory` - The VM host available RAM memory (in MB).
* `local_storage` - The VM host available local storage (in bytes).

