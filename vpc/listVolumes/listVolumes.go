package main

// Lists all volumes on an account, with an optional regional endpoint.

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func main() {
	var apikey string = os.Getenv("IBMCLOUD_API_KEY")
	if apikey == "" {
		log.Fatal("ERROR:\tNo API key set. ($IBMCLOUD_API_KEY)")
	}

	var url = os.Getenv("URL") //This can be empty.

	result := listVolumes(apikey, url)

	j, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(j))
}

func listVolumes(apikey string, url string) (result *vpcv1.VolumeCollection) {
	service, err := vpcv1.NewVpcV1(&vpcv1.VpcV1Options{Authenticator: &core.IamAuthenticator{ApiKey: apikey}, URL: url})
	if err != nil {
		log.Fatal("ERROR:\t", err)
	}

	// returns a Volume Collection, a response header, and error
	result, _, err = service.ListVolumes(&vpcv1.ListVolumesOptions{})
	if err != nil {
		log.Fatal(err)
	}

	return result
}
