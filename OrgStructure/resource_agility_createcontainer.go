package OrgStructure

import (
	"log"
	"os"
	"encoding/json"

	"github.com/csc/csc-agility-terraform-provider-plug-in/api"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/csc/csc-agility-terraform-provider-plug-in/Resource"
)
/*type Container struct{
	XMLName struct{}    `xml:"container"`
	Name 		string `xml:"name"`
	Description	string	`xml:"description"`
}*/

type Container struct {
	XMLName struct{}    `xml:"container"`
	Name  		string 		`xml:"name"`
	HREF 		string   	`xml:"href"`
	Id 			string   	`xml:"id"`
	Rel 		string   	`xml:"rel,omitempty"`
	Type 		string   	`xml:"type,omitempty"`
	Position 	string   	`xml:"position,omitempty"`
}

func resourceCreateSubContainer() *schema.Resource {

	return &schema.Resource{
		Create: CreateSubContainerCreate,
		Read:   CreateSubContainerRead,
		Update: CreateSubContainerUpdate,
		Delete: CreateSubContainerDelete,

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
			"container": &schema.Schema{
				Type:     schema.TypeString,
				Required: 	true,
				ForceNew:	true,
			},
			"containerid": &schema.Schema{
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
type Config struct {
	MaxRetries 	string
	APIURL     	string
	AWSXS  		string
	AWSS  		string
	AWSM  		string
	AWSL  		string
	AWSXL  		string
	AWSXXL 		string
	BCXS    	string
	BCS    		string
	BCM    		string
	BCL    		string
	BCXL    	string
	BCXXL    	string
	AWSCloud 	string
	BizCloud	string
}

var configuration Config
var credentials Resource.ProvCredentials

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

func CreateSubContainerCreate(ResourceData *schema.ResourceData, meta interface{}) error {
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
	//payload := getContainer(containerId,credentials.UserName,credentials.Password)
	api.CreateSubContainer(ResourceData, containerId, credentials.UserName, credentials.Password)
	return nil
}

func CreateSubContainerRead(d *schema.ResourceData, meta interface{}) error {
	// no need to do anything for read state

	return nil
}

func CreateSubContainerUpdate(d *schema.ResourceData, meta interface{}) error {
	// no need to do anything for read state

	return nil
}

func CreateSubContainerDelete(d *schema.ResourceData, meta interface{}) error {
	// no need to do anything for read state

	return nil

}