package OrgStructure

import (
	"log"
	"os"
	"encoding/json"

	"github.com/csc/csc-agility-terraform-provider-plug-in/api"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/csc/csc-agility-terraform-provider-plug-in/Resource"
)

func resourceCreateEnvironments() *schema.Resource {

	return &schema.Resource{
		Create: CreateEnvironmentsCreate,
		Read:   CreateEnvironmentsRead,
		Update: CreateEnvironmentsUpdate,
		Delete: CreateEnvironmentsDelete,

		Schema: map[string]*schema.Schema{
			"parentprojectname": &schema.Schema{
				Type:     schema.TypeString,
				Required: 	true,
				Computed: 	false,
			},
			"parentprojectnameid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: 	true,
				ForceNew:	true,
			},
			"environment": &schema.Schema{
				Type:     schema.TypeString,
				Required: 	true,
				ForceNew:	true,
			},
			"environmentid": &schema.Schema{
				Type:     	schema.TypeString,
				Computed: 	true,
				ForceNew:	true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Required: 	true,
				ForceNew:	true,
			},
			"environmenttype": &schema.Schema{
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

func CreateEnvironmentsCreate(ResourceData *schema.ResourceData, meta interface{}) error {
	credentials = meta.(Resource.ProvCredentials)
	//set up logging
	f, errf := os.OpenFile("./api/agility.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if errf != nil {
		log.Println("error opening file: ", errf)
	}
	defer f.Close()

	log.SetOutput(f)
	projectName := ResourceData.Get("parentprojectname").(string)
	log.Println("the Project name is: ", projectName)
	response, err := api.GetProjectId(string (projectName), credentials.UserName, credentials.Password)
	if err != nil {
		log.Println("Error in getting containerid: ", err)
		return err
	}
	ResourceData.Set("parentprojectnameid",string(response))
	log.Println("the projectid is: ", string(response))
	projectId := ResourceData.Get("parentprojectnameid").(string)
	api.CreateEnvironments(ResourceData, projectId, credentials.UserName, credentials.Password)
	return nil
}

func CreateEnvironmentsRead(d *schema.ResourceData, meta interface{}) error {
	// no need to do anything for read state

	return nil
}

func CreateEnvironmentsUpdate(d *schema.ResourceData, meta interface{}) error {
	// no need to do anything for read state

	return nil
}

func CreateEnvironmentsDelete(d *schema.ResourceData, meta interface{}) error {
	// no need to do anything for read state

	return nil

}