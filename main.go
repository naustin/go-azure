package main

import (
	"context"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"

	"net/http"

	"github.com/gin-gonic/gin"
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

func getResourceTagValue(resourceId *string) (value *string) {
	subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")
	if len(subscriptionID) == 0 {
		log.Panic("AZURE_SUBSCRIPTION_ID is not set.")
	}

	tagName := os.Getenv("AZURE_ENV_TAG_NAME")
	if len(tagName) == 0 {
		log.Panic("AZURE_ENV_TAG_NAME is not set.")
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

	return clientResponse.Tags[tagName]

}

// used for kubernetes liveness probe.  if http 200 not returned then pod will be scheduled for restart.
func livenessProbe(c *gin.Context) {
	c.Status(http.StatusOK)
}

// used for kubernetes readiness probe.  controls traffic to pod.  returning http 200 indicates pod is ready fo rtraffic
func readinessProbe(c *gin.Context) {
	// TODO: check that required environment vars have values
	// TODO: check connection to Azure

	c.Status(http.StatusOK)
}

func RouteBySignalType(c *gin.Context) {

	
	requestRaw, err := c.GetRawData()
	if err != nil {
		log.Printf("Unable to read request, %s", err)
		c.Status(http.StatusBadRequest)
	}

	requestJson := string(requestRaw)

	requestMap := RequestJsonToMap(requestJson)

	alertType := GetAlertType(requestMap)

	if alertType == "Metric" {
		
		
		c.Status(http.StatusOK)
		return
	}

	if alertType == "ServiceHealth" {

		c.Status(http.StatusOK)
		return
	}

	if alertType == "ResourceHealth" {

		c.Status(http.StatusOK)
		return
	}

	log.Printf("No message handler found for this alert type, %s", alertType)
	c.Status(http.StatusBadRequest)
}

func main() {

	//resourceId := "/subscriptions/34558c2d-e4d3-4b3c-91e8-96b795831a5d/resourceGroups/DefaultResourceGroup-EUS"

	//environmentName := *getResourceTagValue(&resourceId)

	//log.Printf("Environment: %s", environmentName)

	router := gin.Default()
	gin.DisableConsoleColor()
	gin.DefaultWriter = io.MultiWriter(os.Stdout)
	router.GET("/liveness-probe", livenessProbe)
	router.GET("/readiness-probe", readinessProbe)
	router.POST("/process-alert", processAlert)

	router.Run("localhost:8080")
}
