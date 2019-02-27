variable "agility_userid" {}
variable "agility_password" {}
provider "agility" {
  userid = "${var.agility_userid}"
  password = "${var.agility_password}"
}


#checkin
resource "agility_checkin" "checkinterraform"{
  locationassetname="Payroll"
  locationassetparentcontainername="HR"
  locationasset="project"
  headversionallowed="true"
  assetname="testpackage"
  asset="package"
  projectname="AgilityFactory"
  /* provisioner "local-exec" {
    command     = "Start-Sleep 10"
    interpreter = ["PowerShell", "-Command"]
  }*/
}

#wait
/*
resource "null_resource" "delay" {
  provisioner "local-exec" {
    command     = "Start-Sleep 15"
    interpreter = ["PowerShell", "-Command"]
  }
  depends_on = ["agility_checkin.checkinterraform"]
}
*/

#approve
/*
resource "agility_approve"  "approve"{
  projectname="${var.project_name}"
  assetname="${var.asset_name}"
  asset="${var.asset_type}"
  state="${var.state}"
  comment="${var.comments}"
  #depends_on = ["null_resource.delay"]
  depends_on = ["agility_checkin.checkinterraform"]

}
*/