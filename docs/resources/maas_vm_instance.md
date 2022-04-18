
# Resource: maas_vm_instance

Provides a resource to create and deploy VMs in MAAS, based on the specified parameters. If no parameters are given, a random machine will be created and deployed using the defaults.

## Example Usage

```terraform
data "maas_vm_hosts" "vm_hosts" {
  id = "rack-01"
}
# single instance
resource "maas_vm_instance" "vm" {
  kvm_no = data.maas_vm_hosts.vm-hosts.recommended
}
# multiply instances
resource "maas_vm_instance" "vms" {
  count = length(data.maas_vm_hosts.vm-hosts.no) * 4
  kvm_no = data.maas_vm_hosts.vm-hosts.no[count.index % length(data.maas_vm_hosts.vm-hosts.no)]
  hostname = "vm-${format("%02d", count.index + 10)}" 
  description = "student ${format("%02d", count.index + 1)}"     
}
```

## Argument Reference

The following arguments are supported:

* `cpu_count` - (Optional) The number of cores used to allocate the MAAS machine.
* `memory` - (Optional) The RAM memory size (in MB) used to allocate the MAAS machine.
* `hostname` - (Optional) The hostname of the MAAS machine to be allocated.
* `zone` - (Optional) The zone name of the MAAS machine to be allocated.
* `pool` - (Optional) The pool name of the MAAS machine to be allocated.
* `tags` - (Optional) A set of tag names that must be assigned on the MAAS machine to be allocated.
* `user_data` - (Optional) Cloud-init user data script that gets run on the machine once it has deployed. A good practice is to set this with `file("/tmp/user-data.txt")`, where `/tmp/user-data.txt` is a cloud-init script.
* `description` - (Optional) Description

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The deployed MAAS machine system ID.
* `fqdn` - The deployed MAAS machine FQDN.
* `hostname` - The deployed MAAS machine hostname.
* `zone` - The deployed MAAS machine zone name.
* `pool` - The deployed MAAS machine pool name.
* `tags` - A set of tag names associated to the deployed MAAS machine.
* `cpu_count` - The number of CPU cores of the deployed MAAS machine.
* `memory` -  The RAM memory size (in GiB) of the deployed MAAS machine.
* `ip_addresses` - A set of IP addressed assigned to the deployed MAAS machine.
