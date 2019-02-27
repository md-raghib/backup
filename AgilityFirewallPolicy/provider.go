package AgilityFirewallPolicy

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/csc/csc-agility-terraform-provider-plug-in/Resource"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {

	// The username and password are retrieved by terraform from the environment variables TF_VAR_agility_userid and TF_VAR_agility_password
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"userid": &schema.Schema{
				Type:     schema.TypeString,
				Required: 	true,
				Computed: 	false,
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Required: 	true,
				Computed: 	false,
			},
		},

		//define the supported resources and point to their respective .go classes
		ResourcesMap: map[string]*schema.Resource{
			
			"agility_firewall":						resourceFirewall(),
			
		},

		ConfigureFunc: providerConfigure,
	}
}

//var descriptions map[string]string

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	creds := Resource.ProvCredentials{
		UserName:        d.Get("userid").(string),
		Password:        d.Get("password").(string),
	}
	return creds, nil
}
