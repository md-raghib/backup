
provider "agility" {
    userid = "${var.userid}"
    password = "${var.password}"
}

# Create a new Linux instance on a small server
resource "agility_license" "test"{
}

