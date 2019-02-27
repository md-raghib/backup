variable "vsphere_user" {
  description = "The User ID for Vsphere."
  default = "vp username"
}
variable "vsphere_password" {
    description = "Password for Vsphere."
    default = "vp pwd"
} 
variable "vsphere_host" {
    description = "The Host address for Vsphere"
    default = "vp host"
}
variable "vsphere_datacenter" {
    description = "The Datacenter Name"
    default = "vp datacenter"
}
variable "vsphere_datastore" {
    description = "The Datastore Name"
    default = "vp datastore"
}
variable "vsphere_network" {
    description = "The Vsphere Network Name"
    default = "v network"
}
variable "plugin_install" {
    description = "The Vsphere Network Name"
    default = "yes install"
}

variable "agility_username" {
    description = "Username of the Agility Appliance"
  default     = "asss"
}	
variable "agility_password" {
    description = "Password of the Agility Appliance"
  default     = "sdfsf"
}	

