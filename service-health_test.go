package main

import (
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHappyPath(t *testing.T) {

	sampleSH := ServiceHealth{
		Environment:          "dev",
		Title:                "API Latency Exceeded Threshold",
		Communication:        "Key vault latency communication.",
		IncidentType:         "",
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

	/*
	   want := regexp.MustCompile(`\b`+name+`\b`)
	   msg, err := Hello("Gladys")
	   if !want.MatchString(msg) || err != nil {
	       t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	   }
	*/
}
