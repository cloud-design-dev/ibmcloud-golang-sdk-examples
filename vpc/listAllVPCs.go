package main

// Lists all VPCs on an account, with an optional resource group and regional endpoint.

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
	var resource_group = os.Getenv("RESOURCE_GROUP") //This can be empty.
	var url = os.Getenv("URL")                       //This can be empty.

	result := listAllVPCs(apikey, &resource_group, url)

	j, _ := json.MarshalIndent(result.Vpcs, "", "  ")
	fmt.Println(string(j))
}

func listAllVPCs(apikey string, resource_group *string, url string) (result *vpcv1.VPCCollection) {
	service, err := vpcv1.NewVpcV1(&vpcv1.VpcV1Options{Authenticator: &core.IamAuthenticator{ApiKey: apikey}, URL: url})
	if err != nil {
		log.Fatal("ERROR:\t", err)
	}

	// returns a VPC Collection, a response header, and error
	result, _, err = service.ListVpcs(&vpcv1.ListVpcsOptions{ResourceGroupID: resource_group})
	if err != nil {
		log.Fatal(err)
	}

	return result
}
