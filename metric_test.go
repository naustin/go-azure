package main

import (
	"testing"
)

func TestMetricAllFields(t *testing.T) {
	sampleMetric := Metric{
		Environment:      "dev",
		AlertRule:        "Alert rule here.",
		MonitorCondition: "Text for monitor condition here.",
		Description:      "Description of alert.",
		EventTimestamp:   "10/3/2024",
		AzurePortalUrl:   "portal.azure.us/#@optumserve-gov.com",
		TargetIds: []string{
			"targetID1",
			"targetID2"},
	}

	output := CreateMetricHtml(sampleMetric)

	if output == "" {
		t.Fatal("No html generated")
	}
	t.Log(output)
}
