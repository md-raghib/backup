//variable "agility_userid" {}
//variable "agility_password" {}
provider "agility" {
  userid = "${var.userid}"
  password = "${var.password}"
}



#Assign Policy

resource "agility_assignpolicy" "assignpolicyterraform"{
  projectname="${var.project_name}"
  policyname="${var.policy_name}"
 //depends_on = ["agility_firewall.terraformfirewall"]
}



