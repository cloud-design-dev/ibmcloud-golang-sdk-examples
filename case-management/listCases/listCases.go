package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/casemanagementv1"
)

func main() {
	var apikey string = os.Getenv("IBMCLOUD_API_KEY")
	if apikey == "" {
		log.Fatal("ERROR:\tNo API key set. ($IBMCLOUD_API_KEY)")
	}
	result, err := listCases(apikey)
	if err != nil {
		log.Fatal(err)
	}

	j, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(j))

}

func listCases(apikey string) (*casemanagementv1.CaseList, error) {
	caseManagementServiceOptions := &casemanagementv1.CaseManagementV1Options{Authenticator: &core.IamAuthenticator{ApiKey: apikey}}
	caseManagementService, err := casemanagementv1.NewCaseManagementV1UsingExternalConfig(caseManagementServiceOptions)
	if err != nil {
		return nil, err
	}

	getCasesOptions := &casemanagementv1.GetCasesOptions{
		Limit:  core.Int64Ptr(int64(10)),
		Fields: []string{"number", "short_description", "severity", "created_at"},
	}

	result, _, err := caseManagementService.GetCases(getCasesOptions)
	if err != nil {
		return nil, err
	}

	return result, nil
}
