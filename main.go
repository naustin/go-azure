package main

import (
	"context"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"

	//"net/http"
    //"github.com/gin-gonic/gin"
)

func getCloudTypeFromEnvVar() (cloudtype cloud.Configuration) {
	// Retrieve the value of the environment variable
	envVarName := "AZURE_CLOUD_TYPE"
	envVarValue := os.Getenv(envVarName)
	azureGovName := "AzureGovernment"
	azureCommercialName := "AzureCommercial"

	// Check if the envVar is valid
	if envVarValue != azureGovName && envVarValue != azureCommercialName {
		log.Panicf("Invalid value for %s: %s. Allowed values are %s or %s", envVarName, envVarValue, azureGovName, azureCommercialName)
	}

	if envVarValue == azureCommercialName {
		return cloud.AzurePublic
	}

	if envVarValue == azureGovName {
		return cloud.AzureGovernment
	}

	log.Panicf("No match found for %s env var. Current value is : %s", envVarName, envVarValue)

	return

}

func getResourceTagValue(resourceId *string, tagName *string) (value *string){
	subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")
	if len(subscriptionID) == 0 {
		log.Panic("AZURE_SUBSCRIPTION_ID is not set.")
	}

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Panic(err)
	}

	opts := azcore.ClientOptions{Cloud: getCloudTypeFromEnvVar()}
	clientFactory, _ := armresources.NewClientFactory(subscriptionID, cred, &arm.ClientOptions{
		ClientOptions: opts,
	})

	resourcesClient := clientFactory.NewClient()

	// Generated from API version 2021-04-01
	clientResponse, err := resourcesClient.GetByID(context.Background(), *resourceId, "2021-04-01", nil)
	if err != nil {
		log.Panic(err)
	}

	return clientResponse.Tags[*tagName]

}

func main() {

	resourceId := "/subscriptions/34558c2d-e4d3-4b3c-91e8-96b795831a5d/resourceGroups/DefaultResourceGroup-EUS"
	tagName := "Environment"

	environmentName := *getResourceTagValue(&resourceId, &tagName)

	log.Printf("Environment: %s", environmentName)
}
