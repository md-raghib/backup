variable "userid" {
    description="The Username of the Agility Platform"
    default = "admin"
}

variable "password" {
    description="The Username of the Agility Platform"
    default = "M3sh@dmin!"}

variable "project_name" {
  default = "Agility Factory"}
 

variable "firewall_name" {
  default = "CollectorFirewall"}

variable "firewall_desc" {
  default = "Firewall Policy"}

variable "direction" {
  default = "Input"}

variable "protocol_prefix" {
  default = "0.0.0.0/0"}

variable "protocol_allowed" {
  default = "true"}

variable "firewall_type" {
  default = "Firewall"}

variable "protocol_name1" {
  default = "https"}
variable "protocol_minport1" {
  default = "8443"}

variable "protocol_maxport1" {
  default = "8443"}

variable "protocol1" {
  default = "tcp"}
  
  
variable "protocol_name2" {
  default = "RealTimeMonitoring"}

variable "protocol_minport2" {
  default = "2187"}

variable "protocol_maxport2" {
  default = "2187"}

variable "protocol2" {
  default = "tcp"}
  
variable "protocol_name3" {
  default = "AgilityMonitorTcp"}

variable "protocol_minport3" {
  default = "8649"}

variable "protocol_maxport3" {
  default = "8649"}
  
variable "protocol3" {
  default = "tcp"}  
  
variable "protocol_name4" {
  default = "AgilityMonitorUdp"}

variable "protocol_minport4" {
  default = "8649"}

variable "protocol_maxport4" {
  default = "8649"}
  
variable "protocol4" {
  default = "udp"}

variable "blueprint_name" {
    default = "varblueprint_name"}

variable "blueprint_desc" {
    default = "varblueprint_description"}

variable "package_name" {
    default = "varpackage_name"}

variable "stack_name" {
  default = "varstack_name"}

variable "policy_name" {
    default = "varpolicy_name"}

variable "headversion_allow" {
    default = "varheadversion_allow"}

variable "workload_name" {
    default = "varworkload_name"}

variable "policyassignment_name" {
  default = "varpolicyassignment_name"}

variable "product_name" {
    default = "varproduct_name"}

variable "product_desc" {
    default = "varproduct_desc"}

variable "product_type" {
  default = "varproduct_type"}

variable "item_type" {
    default = "varitem_type"}

variable "item_name" {
    default = "varitem_name"}

variable "category" {
    default = "varcategory"}

variable "operating_system" {
  default = "varoperating_system"}

variable "container_name" {
  default = "varcontainer_name"}

variable "headversion_allowed" {
  default = "varheadversion_allowed"}

variable "asset_name" {
  default = "varasset_name"}

variable "asset_type" {
  default = "varasset_type"}

variable "state" {
  default = "varstate"}

variable "comments" {
  default = "varcomments"}
