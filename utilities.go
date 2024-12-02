package main

import (
	"bytes"
	"encoding/json"
	"html/template"
	"time"
	"os"
	"log"
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/html"
)

type CommonAlertSchema struct {
	Headers struct {
		Expect               string `json:"Expect"`
		Host                 string `json:"Host"`
		MaxForwards          string `json:"Max-Forwards"`
		UserAgent            string `json:"User-Agent"`
		XCorrelationContext  string `json:"X-CorrelationContext"`
		XARRLOGID            string `json:"X-ARR-LOG-ID"`
		CLIENTIP             string `json:"CLIENT-IP"`
		DISGUISEDHOST        string `json:"DISGUISED-HOST"`
		XSITEDEPLOYMENTID    string `json:"X-SITE-DEPLOYMENT-ID"`
		WASDEFAULTHOSTNAME   string `json:"WAS-DEFAULT-HOSTNAME"`
		XForwardedProto      string `json:"X-Forwarded-Proto"`
		XAppServiceProto     string `json:"X-AppService-Proto"`
		XARRSSL              string `json:"X-ARR-SSL"`
		XForwardedTLSVersion string `json:"X-Forwarded-TlsVersion"`
		XForwardedFor        string `json:"X-Forwarded-For"`
		XOriginalURL         string `json:"X-Original-URL"`
		XWAWSUnencodedURL    string `json:"X-WAWS-Unencoded-URL"`
		ContentLength        string `json:"Content-Length"`
		ContentType          string `json:"Content-Type"`
	} `json:"headers"`
	Body struct {
		SchemaID string `json:"schemaId"`
		Data     struct {
			Essentials struct {
				AlertID             string    `json:"alertId"`
				AlertRule           string    `json:"alertRule"`
				TargetResourceType  string    `json:"targetResourceType"`
				AlertRuleID         string    `json:"alertRuleID"`
				Severity            string    `json:"severity"`
				SignalType          string    `json:"signalType"`
				MonitorCondition    string    `json:"monitorCondition"`
				TargetResourceGroup string    `json:"targetResourceGroup"`
				MonitoringService   string    `json:"monitoringService"`
				AlertTargetIDs      []string  `json:"alertTargetIDs"`
				ConfigurationItems  []string  `json:"configurationItems"`
				OriginAlertID       string    `json:"originAlertId"`
				FiredDateTime       time.Time `json:"firedDateTime"`
				ResolvedDateTime    time.Time `json:"resolvedDateTime"`
				Description         string    `json:"description"`
				EssentialsVersion   string    `json:"essentialsVersion"`
				AlertContextVersion string    `json:"alertContextVersion"`
			} `json:"essentials"`
			AlertContext struct {
				Properties struct {
					From string `json:"from"`
				} `json:"properties"`
				ConditionType string `json:"conditionType"`
				Condition     struct {
					WindowSize string `json:"windowSize"`
					AllOf      []struct {
						MetricName      string        `json:"metricName"`
						MetricNamespace string        `json:"metricNamespace"`
						Operator        string        `json:"operator"`
						Threshold       string        `json:"threshold"`
						TimeAggregation string        `json:"timeAggregation"`
						Dimensions      []interface{} `json:"dimensions"`
						MetricValue     int           `json:"metricValue"`
						WebTestName     interface{}   `json:"webTestName"`
					} `json:"allOf"`
					StaticThresholdFailingPeriods struct {
						NumberOfEvaluationPeriods int `json:"numberOfEvaluationPeriods"`
						MinFailingPeriodsToAlert  int `json:"minFailingPeriodsToAlert"`
					} `json:"staticThresholdFailingPeriods"`
					WindowStartTime time.Time `json:"windowStartTime"`
					WindowEndTime   time.Time `json:"windowEndTime"`
				} `json:"condition"`
			} `json:"alertContext"`
			CustomProperties struct {
				From string `json:"from"`
			} `json:"customProperties"`
		} `json:"data"`
	} `json:"body"`
}


func HtmlMinify(content string) string {

	m := minify.New()
	m.AddFunc("text/html", html.Minify)

	s, err := m.String("text/html", content)
	if err != nil {
		panic(err)
	}

	return s
}

func BuildHtmlFromTemplate(tmpl string, data any) string {

	t, err := template.New("html").Parse(tmpl)
	if err != nil {
		panic(err)
	}
	var buf bytes.Buffer
	err = t.Execute(&buf, data)
	if err != nil {
		panic(err)
	}

	return buf.String()
}

func RequestJsonToMap(jsonString string) CommonAlertSchema {
	
	// Parse the JSON string
	var req CommonAlertSchema
	err := json.Unmarshal([]byte(jsonString), &req)
	if err != nil {
		panic(err)
	}

	return req
}

func GetAlertType(reqMap CommonAlertSchema) string {
	return reqMap.Body.Data.Essentials.SignalType
}

func GetResourceTagValue(resourceId *string) (value *string) {
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