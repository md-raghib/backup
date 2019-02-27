package OrgStructure

import (
	"log"
	"os"
	"encoding/json"

	"github.com/csc/csc-agility-terraform-provider-plug-in/api"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/csc/csc-agility-terraform-provider-plug-in/Resource"
)

func resourceCreateSubProject() *schema.Resource {

	return &schema.Resource{
		Create: CreateSubProjectCreate,
		Read:   CreateSubProjectRead,
		Update: CreateSubProjectUpdate,
		Delete: CreateSubProjectDelete,

		Schema: map[string]*schema.Schema{
			"parentcontainername": &schema.Schema{
				Type:     schema.TypeString,
				Required: 	true,
				Computed: 	false,
			},
			"parentcontainernameid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: 	true,
				ForceNew:	true,
			},
			"project": &schema.Schema{
				Type:     schema.TypeString,
				Required: 	true,
				ForceNew:	true,
			},
			"projectid": &schema.Schema{
				Type:     	schema.TypeString,
				Computed: 	true,
				ForceNew:	true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Required: 	true,
				ForceNew:	true,
			},
		},
	}
}

func init(){
	file, err1 := os.Open("./api/conf.json")
	if err1 != nil {
		log.Println("error:", err1)
	}
	decoder := json.NewDecoder(file)
	configuration = Config{}
	err := decoder.Decode(&configuration)
	if err != nil {
		log.Println("error:", err)
	}

	/*err2 := file.Close()
	log.Printf("err2: %v\n", err2)*/
}

func CreateSubProjectCreate(ResourceData *schema.ResourceData, meta interface{}) error {
	credentials = meta.(Resource.ProvCredentials)
	//set up logging
	f, errf := os.OpenFile("./api/agility.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if errf != nil {
		log.Println("error opening file: ", errf)
	}
	defer f.Close()

	log.SetOutput(f)
	containerName := ResourceData.Get("parentcontainername").(string)
	log.Println("the Container name is: ", containerName)
	response, err := api.GetContainerId(string (containerName), credentials.UserName, credentials.Password)
	if err != nil {
		log.Println("Error in getting containerid: ", err)
		return err
	}
	ResourceData.Set("parentcontainernameid",string(response))
	log.Println("the containerid is: ", string(response))
	containerId := ResourceData.Get("parentcontainernameid").(string)
	api.CreateSubProject(ResourceData, containerId, credentials.UserName, credentials.Password)
	return nil
}

func CreateSubProjectRead(d *schema.ResourceData, meta interface{}) error {
	// no need to do anything for read state

	return nil
}

func CreateSubProjectUpdate(d *schema.ResourceData, meta interface{}) error {
	// no need to do anything for read state

	return nil
}

func CreateSubProjectDelete(d *schema.ResourceData, meta interface{}) error {
	// no need to do anything for read state

	return nil

}