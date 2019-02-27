
provider "agility" {
    userid = "${var.userid}"
    password = "${var.password}"
}

resource "agility_createcontainer" "Container" {
  parentcontainername = "${var.headcontainername}"
  container = "${var.parentcontainer3}"
  description = "Container created via terraform"
}
/*
resource "agility_createproject" "Project"{
  parentcontainername="${var.parentcontainer3}"
  project="${var.project2}"
  description="Project created via terraform"
  depends_on = ["agility_createcontainer.Container"]
}

resource "agility_createenvironment" "Environment" {
  parentprojectname = "${var.project2}"
  environment = "${var.environment1}"
  description = "Environment created via terraform"
  environmenttype = "${var.environment1_type}"
  depends_on = ["agility_createproject.Project"]
}
*/