package AddCloud

import (
	"log"
	"os"
	"encoding/json"

	"github.com/csc/csc-agility-terraform-provider-plug-in/api"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/csc/csc-agility-terraform-provider-plug-in/Resource"
)
func resourceAddCloudProvider() *schema.Resource {

	return &schema.Resource{
		Create: AddCloudProviderCreate,
		Read:   AddCloudProviderRead,
		Update: AddCloudProviderUpdate,
		Delete: AddCloudProviderDelete,

		Schema: map[string]*schema.Schema{
			"cloudname": &schema.Schema{
				Type:     schema.TypeString,
				Required: 	true,
				Computed: 	false,
			},
			/*"updatedcloudname": &schema.Schema{
				Type:     schema.TypeString,
				Required: 	true,
				Computed: 	false,
			},*/
			"cloudid": &schema.Schema{
				Type: 	schema.TypeString,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Required: 	true,
				ForceNew:	true,
			},
			"cloudtype": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew:	true,
			},
			/*"cloudtypeid": &schema.Schema{
				Type: 	schema.TypeString,
				Computed: true,
			},*/
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Required: 	true,
				ForceNew:	true,
			},
			"publickey": &schema.Schema{
				Type:     schema.TypeString,
				Required: 	true,
				ForceNew:	true,
			},
			"privatekey": &schema.Schema{
				Type:     schema.TypeString,
				Required: 	true,
				ForceNew:	true,
			},
			"awsaccountnumber": &schema.Schema{
				Type:     schema.TypeString,
				Required: 	true,
				ForceNew:	true,
			},
			/*"operation": &schema.Schema{
				Type:     schema.TypeString,
				Required:   true,
				Computed:   false,
			},*/
			"credentialtype": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
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

func AddCloudProviderCreate(ResourceData *schema.ResourceData, meta interface{}) error {
	credentials = meta.(Resource.ProvCredentials)
	//set up logging
	f, errf := os.OpenFile("./api/agility.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if errf != nil {
		log.Println("error opening file: ", errf)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println("!!!Inside create!!!")
	api.AddCloudProvider(ResourceData, credentials.UserName, credentials.Password)

	return nil
}

func AddCloudProviderRead(d *schema.ResourceData, meta interface{}) error {
	// no need to do anything for read state

	return nil
}

func AddCloudProviderUpdate(d *schema.ResourceData, meta interface{}) error {
	// no need to do anything for read state

	return nil
}

func AddCloudProviderDelete(d *schema.ResourceData, meta interface{}) error {
	// no need to do anything for read state

	return nil

}