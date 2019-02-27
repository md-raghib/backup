//variable "agility_userid" {}
//variable "agility_password" {}
provider "agility" {
  userid = "${var.userid}"
  password = "${var.password}"
}

#as best practice we are publishing from root container

resource "agility_publish" "terraformblueprint" {
  productname = "${var.product_name}"
  productdesc = "${var.product_desc}"
  producttype = "${var.product_type}"
  itemtype = "${var.item_type}"
  itemname = "${var.item_name}"
  category = "${var.category}"
  operatingsystem = "${var.operating_system}"
}


/*
resource "agility_headversion" "version" {
  asset="script"
  projectname="Agility Factory"
}
*/

#Script ID

/*resource "agility_scriptId" "samplescriptid"{
  scriptname="jmd"
  projectname="Agility Factory"
}
*/