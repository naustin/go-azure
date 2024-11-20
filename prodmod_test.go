package main

import (
	"testing"
)

func TestProdModAllFields(t *testing.T) {
	sampleProdMod := ProdMod{
		Environment:         "dev",
		EventTimestamp: "10/3/2024",
		TargetIds: []string{
			"targetID1",
			"targetID2"},
	}

	output := CreateProdModHtml(sampleProdMod)

	if output == "" {
		t.Fatal("No html generated")
	}
	t.Log(output)
}