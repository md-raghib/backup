package Resource

import (
	"fmt"
	"log"
	"os"

	"github.com/csc/csc-agility-terraform-provider-plug-in/api"
	"github.com/hashicorp/terraform/helper/schema"
)

//Deleted as no longer used as a Terraform Resource
/*func resourceAgilityEnvironment() *schema.Resource {

	return &schema.Resource{
		Create: resourceAgilityEnvironmentCreate,
		Read:   resourceAgilityEnvironmentRead,
		//Update: resourceAgilityEnvironmentUpdate,
		Delete: resourceAgilityEnvironmentDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: 	true,
				ForceNew:	true,
			},
			"projectid": &schema.Schema{
				Type:     	schema.TypeString,
				Required:	true,
				ForceNew:	true,
			},
		},
	}
}*/

// func resourceAgilityEnvironmentCreate(d *schema.ResourceData, meta interface{}) error {
// 	// set up logging
// 	f, errf := os.OpenFile("agility.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
//     if errf != nil {
//         log.Println("error opening file: ", errf)
//     }
//     defer f.Close()

//     log.SetOutput(f)

// 	// Associate to Project if specified, which it should be
// 	// if the .tf file is configured correctly
// 	projectid, ok_projectid := d.GetOk("projectid")
// 	log.Println("the Project Id is: ", projectid)

// 	if ok_projectid {
// 		//get the ID of the Environment by calling the Agility API
// 		response, err := api.Getenvironmentid(d.Get("name").(string), projectid.(string))
// 		if err != nil {
// 			return err
// 		}

// 		log.Println("the Environment Id is: ", response)
// 		//set the ID as the ID of this resource
// 		d.SetId(string(response))

// 		return nil
// 	} else {
// 		return fmt.Errorf("No projectid was provided")
// 	}

// 	return nil
// }

func checkEnvironment(d *schema.ResourceData) error {
	// set up logging
	f, errf := os.OpenFile("agility.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
    if errf != nil {
        log.Println("error opening file: ", errf)
    }
    defer f.Close()

    log.SetOutput(f)

	// Associate to Project if specified, which it should be
	// if the .tf file is configured correctly
	projectid, ok_projectid := d.GetOk("projectid")
	log.Println("the Project Id is: ", projectid)

	if ok_projectid {
		//get the ID of the Environment by calling the Agility API
		response, err := api.GetEnvironmentId(d.Get("environment").(string), projectid.(string), credentials.UserName, credentials.Password)
		if err != nil {
			return err
		}

		log.Println("the Environment Id is: ", response)
		//set the ID as the ID of this resource
		d.Set("environmentid",string(response))

		return nil
	} else {
		return fmt.Errorf("No projectid was provided")
	}

	return nil
}

//Deleted as no longer used as a Terraform Resource
// func resourceAgilityEnvironmentRead(d *schema.ResourceData, meta interface{}) error {
// 	// Associate to Project if specified, which it should be
// 	// if the .tf file is configured correctly
// 	projectid, ok_projectid := d.GetOk("projectid")

// 	if ok_projectid {
// 		//get the ID of the Environment by calling the Agility API
// 		response, err := api.Getenvironmentid(d.Get("environment").(string), projectid.(string))
// 		if err != nil {
// 			return err
// 		}

// 		//set the ID as the ID of this resource
// 		d.SetId(string(response))

// 		return nil
// 	} else {
// 		return fmt.Errorf("No projectid was provided")
// 	}

// 	return nil
// }

// func resourceAgilityEnvironmentUpdate(d *schema.ResourceData, meta interface{}) error {
// 	// Associate to Project if specified, which it should be
// 	// if the .tf file is configured correctly
// 	projectid, ok_projectid := d.GetOk("projectid")

// 	if ok_projectid {
// 		//get the ID of the Environment by calling the Agility API
// 		response, err := api.Getenvironmentid(d.Get("environment").(string), projectid.(string))
// 		if err != nil {
// 			return err
// 		}

// 		//set the ID as the ID of this resource
// 		d.SetId(string(response))

// 		return nil
// 	} else {
// 		return fmt.Errorf("No projectid was provided")
// 	}

// 	return nil
// }

// func resourceAgilityEnvironmentDelete(d *schema.ResourceData, meta interface{}) error {
// 	// we don't delete the environment in agility, so just remove the resource from Terraform
// 	// by removing the ID
// 	d.SetId("")

// 	return nil
// }
