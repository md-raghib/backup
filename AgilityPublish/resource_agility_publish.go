package AgilityPublish

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
    Name        string `xml:"name"`
    Description string  `xml:"description"`
}*/

func resourceAgilityPublish() *schema.Resource {

	return &schema.Resource{
		Create: resourcePublishCreate,
		Read:   resourcePublishRead,
		Update: resourcePublishUpdate,
		Delete: resourcePublishDelete,

		Schema: map[string]*schema.Schema{
			"productname": &schema.Schema{
				Type:     schema.TypeString,
				Required:   true,
				Computed:   false,
			},
			"productdesc": &schema.Schema{
				Type:     schema.TypeString,
				Required:   true,
				Computed:   false,
			},
			"producttype": &schema.Schema{
				Type:     schema.TypeString,
				Required:   true,
				Computed:   false,
			},
			"itemtype": &schema.Schema{
				Type:     schema.TypeString,
				Required:   true,
				Computed:   false,
			},
			"itemname": &schema.Schema{
				Type:     schema.TypeString,
				Required:   true,
				Computed:   false,
			},
			"category": &schema.Schema{
				Type:     schema.TypeString,
				Required:   true,
				Computed:   false,
			},
			"operatingsystem": &schema.Schema{
				Type:     schema.TypeString,
				Required:   true,
				ForceNew:   true,
			},
			"assetid": &schema.Schema{
				Type:       schema.TypeString,
				Computed:   true,
				ForceNew:   true,
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

func resourcePublishCreate(ResourceData *schema.ResourceData, meta interface{}) error {
	credentials = meta.(Resource.ProvCredentials)
	//set up logging
	f, errf := os.OpenFile("./api/agility.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if errf != nil {
		log.Println("error opening file: ", errf)
	}
	defer f.Close()

	log.SetOutput(f)

	api.Publish(ResourceData,credentials.UserName,credentials.Password)

	return nil
}

func resourcePublishRead(d *schema.ResourceData, meta interface{}) error {
	// no need to do anything for read state

	return nil
}

func resourcePublishUpdate(d *schema.ResourceData, meta interface{}) error {
	// no need to do anything for read state

	return nil
}

func resourcePublishDelete(d *schema.ResourceData, meta interface{}) error {
	// no need to do anything for read state

	return nil
}
