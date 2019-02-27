package AddCloud

import (
	"log"
	"os"
	"encoding/json"

	"github.com/csc/csc-agility-terraform-provider-plug-in/api"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/csc/csc-agility-terraform-provider-plug-in/Resource"
)

func resourceSyncCloudProvider() *schema.Resource {

	return &schema.Resource{
		Create: SyncCloudProviderCreate,
		Read:   SyncCloudProviderRead,
		Update: SyncCloudProviderUpdate,
		Delete: SyncCloudProviderDelete,

		Schema: map[string]*schema.Schema{
			"cloudname": &schema.Schema{
				Type:     schema.TypeString,
				Required: 	true,
				Computed: 	false,
			},
			"cloudid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: 	true,
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

func SyncCloudProviderCreate(ResourceData *schema.ResourceData, meta interface{}) error {
	credentials = meta.(Resource.ProvCredentials)
	//set up logging
	f, errf := os.OpenFile("./api/agility.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if errf != nil {
		log.Println("error opening file: ", errf)
	}
	defer f.Close()

	log.SetOutput(f)
	cloudName := ResourceData.Get("cloudname").(string)
	log.Println("the Cloud name is: ", cloudName)
	response, err := api.GetCloudId(string (cloudName), credentials.UserName, credentials.Password)
	if err != nil {
		log.Println("Error in getting cloudid: ", err)
		return err
	}
	ResourceData.Set("cloudid",string(response))
	log.Println("the cloudid is: ", string(response))
	cloudId := ResourceData.Get("cloudid").(string)
	api.SyncCloudProvider(ResourceData, cloudId, credentials.UserName, credentials.Password)
	return nil
}

func SyncCloudProviderRead(d *schema.ResourceData, meta interface{}) error {
	// no need to do anything for read state

	return nil
}

func SyncCloudProviderUpdate(d *schema.ResourceData, meta interface{}) error {
	// no need to do anything for read state

	return nil
}

func SyncCloudProviderDelete(d *schema.ResourceData, meta interface{}) error {
	// no need to do anything for read state

	return nil

}