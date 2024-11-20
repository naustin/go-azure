package main

import (
	"testing"
)

func TestServiceHealthAllFields(t *testing.T) {

	sampleSH := ServiceHealth{
		Environment:          "dev",
		Title:                "API Latency Exceeded Threshold",
		Communication:        "Key vault latency communication.",
		IncidentType:         "Incident type here.",
		ImpactStartTime:      "10/22/2024",
		ImpactMitigationTime: "10/23/2024",
		TargetIds: []string{
			"targetID1",
			"targetID2"},
		ImpactedServices: []string{
			"service1",
			"service2",
			"service3"},
	}

	output := CreateServiceHealthHtml(sampleSH)

	if output == "" {
		t.Fatal("No html generated")
	}
	t.Log(output)

}

func TestOnlyServiceHealthEnvAndTitle(t *testing.T) {

	sampleSH := ServiceHealth{
		Environment:          "dev",
		Title:                "API Latency Exceeded Threshold",
		Communication:        "",
		IncidentType:         "",
		ImpactStartTime:      "",
		ImpactMitigationTime: "",
		TargetIds:            nil,
		ImpactedServices:     nil,
	}

	output := CreateServiceHealthHtml(sampleSH)

	if output == "" {
		t.Fatal("No html generated")
	}
	t.Log(output)

}
