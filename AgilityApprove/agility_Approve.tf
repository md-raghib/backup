//variable "agility_userid" {}
//variable "agility_password" {}
provider "agility" {
  userid = "${var.userid}"
  password = "${var.password}"
}


#approve

resource "agility_approve"  "approve"{
  projectname="${var.project_name}"
  assetname="${var.asset_name}"
  asset="${var.asset_type}"
  state="${var.state}"
  comment="${var.comments}"
  //depends_on = ["agility_checkin.checkinterraform"]
}


