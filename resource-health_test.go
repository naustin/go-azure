package main

import (
	"testing"
)

func TestResourceHealthAllFields(t *testing.T) {

	sampleRH := ResourceHealth{
		Environment:         "dev",
		Title:               "API Latency Exceeded Threshold",
		CurrentHealthStatus: "Unhealthy",
		EventTimestamp:      "10/22/2024",
		Level:               "Level High",
		Issue:               "Issue text here",
		Type:                "Alert type resource health",
		Cause:               "Uknown cause",
		AzurePortalUrl:      "portal.azure.us/#@optumserve-gov.com",
		TargetIds: []string{
			"targetID1",
			"targetID2"},
	}

	output := CreateResourceHealthHtml(sampleRH)

	if output == "" {
		t.Fatal("No html generated")
	}
	t.Log(output)

}
