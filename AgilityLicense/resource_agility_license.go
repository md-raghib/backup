package AgilityLicense

import (
	"log"
	"strings"
	"os"
	//"encoding/json"

	"github.com/csc/csc-agility-terraform-provider-plug-in/api"
	"github.com/hashicorp/terraform/helper/schema"

	"encoding/json"

	"github.com/csc/csc-agility-terraform-provider-plug-in/Resource"
)

type SMLicense struct{
	XMLName struct{}    `"SMLicense"`
	XMLNS 		string `"xmlns,attr,omitempty"`
	LicenseVersion 	string	`"LicenseVersion"`
	Licensee		string	`"Licensee"`
	NumberOfClients	string	`"NumberOfClients"`
	MaxPermittedInstances	string `"MaxPermittedInstances"`
	DBQueryType		string	`"DBQueryType"`
	Signature		string 	`"Signature"`
}
type Adapters struct{
	XMLName struct{}    `"Adapters"`
	Adapter string		`"Adapter"`

}
type Modules struct{
	XMLName struct{}    `"Modules"`
	Module string		`"Module"`
}
type IssueDate struct{
	XMLName struct{}	`"IssueDate"`
	IssueDay	string	`"IssueDay"`
	IssueMonth	string	`"IssueMonth"`
	IssueYear	string	`"IssueYear"`
}
type ExpiryDate struct{
	XMLName struct{}	`"ExpiryDate"`
	ExpiryDay	string	`"ExpiryDay"`
	ExpiryMonth	string	`"ExpiryMonth"`
	ExpiryYear	string	`"ExpiryYear"`
}
type ServerNodeLock struct{
	XMLName struct{}	`"ServerNodeLock"`
	Mask	string	`"Mask"`
	IPAddr	string	`"IPAddr"`
}

func resourceLicenseUpload() *schema.Resource {

	return &schema.Resource{
		Create: LicenseUpload,
		Read: 	LicenseUploadRead,
		Delete:	LicenseUploadDelete,
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

func init() {
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
}

func LicenseUpload(ResourceData *schema.ResourceData, meta interface{}) error {
	credentials = meta.(Resource.ProvCredentials)
	f, errf := os.OpenFile("./api/agility.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if errf != nil {
		log.Println("error opening file: ", errf)
	}
	defer f.Close()
	log.SetOutput(f)

	licenseresponse := api.LicenseUpload(credentials.UserName, credentials.Password)
	r := strings.NewReader(string(licenseresponse))
	log.Println("Deploy response is : ", r)

	return nil
}

func LicenseUploadRead(d *schema.ResourceData, meta interface{}) error {
	// no need to do anything for read state

	return nil
}
func LicenseUploadDelete(d *schema.ResourceData, meta interface{}) error {
	// no need to do anything for read state

	return nil
}
