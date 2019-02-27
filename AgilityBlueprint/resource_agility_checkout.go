package AgilityBlueprint

import (
	"log"
	"os"
	"encoding/json"

	"github.com/csc/csc-agility-terraform-provider-plug-in/api"
	"github.com/csc/csc-agility-terraform-provider-plug-in/Resource"
	"github.com/hashicorp/terraform/helper/schema"
)
/*type Container struct{
    XMLName struct{}    `xml:"container"`
    Name        string `xml:"name"`
    Description string  `xml:"description"`
}*/

func resourceAgilityCheckOut() *schema.Resource {

	return &schema.Resource{
		Create: resourceCheckoutCreate,
		Read:   resourceCheckoutRead,
		Update: resourceCheckoutUpdate,
		Delete: resourceCheckoutDelete,

		Schema: map[string]*schema.Schema{
			"projectname": &schema.Schema{
				Type:     schema.TypeString,
				Required:   true,
				Computed:   false,
			},
			"containername": &schema.Schema{
				Type:     schema.TypeString,
				Required:   true,
				Computed:   false,
			},
			"projectid": &schema.Schema{
				Type:       schema.TypeString,
				Computed:   true,
				ForceNew:   true,
			},
			"containerid": &schema.Schema{
				Type:       schema.TypeString,
				Computed:   true,
				ForceNew:   true,
			},
			"assetname": &schema.Schema{
				Type:     schema.TypeString,
				Required:   true,
				Computed:   false,
			},
			"asset": &schema.Schema{
				Type:     schema.TypeString,
				Required:   true,
				Computed:   false,
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

func resourceCheckoutCreate(ResourceData *schema.ResourceData, meta interface{}) error {
	credentials = meta.(Resource.ProvCredentials)
	//set up logging
	f, errf := os.OpenFile("./api/agility.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if errf != nil {
		log.Println("error opening file: ", errf)
	}
	defer f.Close()

	log.SetOutput(f)
	projectName := ResourceData.Get("projectname").(string)
	log.Println("the Project name is: ", projectName)
	response, err := api.GetProjectId(string (projectName), credentials.UserName, credentials.Password)
	if err != nil {
		log.Println("Error in getting projectid: ", err)
		return err
	}
	ResourceData.Set("projectid",string(response))
	log.Println("the projectid is: ", string(response))
	projectId := ResourceData.Get("projectid").(string)

	api.CheckOut(ResourceData,projectId,credentials.UserName,credentials.Password)

	return nil
}

func resourceCheckoutRead(d *schema.ResourceData, meta interface{}) error {
	// no need to do anything for read state

	return nil
}

func resourceCheckoutUpdate(d *schema.ResourceData, meta interface{}) error {
	// no need to do anything for read state

	return nil
}

func resourceCheckoutDelete(d *schema.ResourceData, meta interface{}) error {
	// no need to do anything for read state

	return nil
}

