package main

import (
	"io"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"

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

func RouteBySignalType(c *gin.Context, tagLookupFn ResourceTagFunc) {

	
	requestRaw, err := c.GetRawData()
	if err != nil {
		log.Printf("Unable to read request, %s", err)
		c.Status(http.StatusBadRequest)
	}

	requestJson := string(requestRaw)

	requestMap := RequestJsonToMap(requestJson)

	alertType := GetAlertType(requestMap)

	if alertType == "Metric" {
		
		msg := ProcessMetricRequest(requestMap, tagLookupFn)
		c.Data(http.StatusOK, "application/text", []byte(msg))
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
	Group := router.Group("api/v1/")
    {
	  Group.GET("/liveness-probe", livenessProbe)
	  Group.GET("/readiness-probe", readinessProbe)
	  Group.POST("/process-alert", RouteBySignalType(&gin.Context, GetResourceTagValue) )
	}
	

	router.Run("localhost:8080")
}
