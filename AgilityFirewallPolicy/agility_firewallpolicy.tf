//variable "agility_userid" {}
//variable "agility_password" {}
provider "agility" {
  userid = "${var.userid}"
  password = "${var.password}"
}


#firewall

resource "agility_firewall" "terraformfirewall"{
  projectname = "${var.project_name}"
  firewallname = "${var.firewall_name}"
  firewalldesc="${var.firewall_desc}"
  direction="${var.direction}"
  protocolprefix="${var.protocol_prefix}"
  protocolallowed="${var.protocol_allowed}"
  firewalltype="${var.firewall_type}"

  protocolname1="${var.protocol_name1}"
  protocoldesc1="HTTPS"
  protocolminport1="${var.protocol_minport1}"
  protocolmaxport1="${var.protocol_maxport1}"
  protocol1="${var.protocol1}"

  protocolname2="${var.protocol_name2}"
  protocoldesc2="RealTimeMonitoring"
  protocolminport2="${var.protocol_minport2}"
  protocolmaxport2="${var.protocol_maxport2}"
  protocol2="${var.protocol2}"

  protocolname3="${var.protocol_name3}"
  protocoldesc3="AgilityMonitorTcp"
  protocolminport3="${var.protocol_minport3}"
  protocolmaxport3="${var.protocol_maxport3}"
  protocol3="${var.protocol3}"

  protocolname4="${var.protocol_name4}"
  protocoldesc4="AgilityMonitorUdp"
  protocolminport4="${var.protocol_minport4}"
  protocolmaxport4="${var.protocol_maxport4}"
  protocol4="${var.protocol4}"
 #depends_on = ["agility_createpackage.terraformpackage"]
}


