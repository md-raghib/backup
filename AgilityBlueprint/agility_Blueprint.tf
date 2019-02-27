//variable "agility_userid" {}
//variable "agility_password" {}
provider "agility" {
  userid = "${var.userid}"
  password = "${var.password}"
}

#Assign Policy
/*
resource "agility_assignpolicy" "assignpolicyterraform"{
  projectname="${var.project_name}"
  policyname="${var.policy_name}"
 //depends_on = ["agility_firewall.terraformfirewall"]
}
*/

resource "agility_checkOut" "assetCheckOut"{
  containername="${var.container_name}"
  assetname="${var.package_name}"
  asset="${var.asset_type}"
  projectname="${var.project_name}"
}

#Create Blueprint
resource "agility_blueprint" "createblueprintterraform" {
  projectname = "${var.project_name}"
  blueprintname = "${var.blueprint_name}"
  blueprintdesc = "${var.blueprint_desc}"
  stackname = "${var.stack_name}"
  packagename = "${var.package_name}"
  policyname = "${var.policy_name}"
  headversionallowed = "${var.headversion_allowed}"
  workloadname = "${var.workload_name}"
  policyassignmentname = "${var.policyassignment_name}"
  depends_on = ["agility_checkOut.assetCheckOut"]
}




#unassign
/*
resource "agility_unassignpolicy" "unassignpolicyterraform"{
  projectname="Agility Factory"
  policyname="CollectorFirewallTerraform"
  depends_on = ["agility_blueprint.createblueprintterraform"]
}
*/
