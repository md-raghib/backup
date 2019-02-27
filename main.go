package main

import (
    "github.com/hashicorp/terraform/plugin"
    "github.com/hashicorp/terraform/terraform"


    "github.com/csc/csc-agility-terraform-provider-plug-in/AgilityCheckinAndApprove"
)

func main() {
    plugin.Serve(&plugin.ServeOpts{
        ProviderFunc: func() terraform.ResourceProvider {
            return AgilityCheckinAndApprove.Provider()
        },
    })
}
