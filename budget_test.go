package main

import (
	"testing"
)

func TestBudgetAllFields(t *testing.T) {
	sampleBudget := Budget{
		Environment:              "dev",
		ThresholdType:            "Threshold type.",
		ForecastedTotalForPeriod: "100.00",
		Description:              "Description of alert.",
		AzurePortalUrl:           "portal.azure.us/#@optumserve-gov.com",
		SpentAmount:              "125.00",
		BudgetThreshold:          "105.00",
		BudgetId:                 "ID0002343",
	}

	output := CreateBudgetHtml(sampleBudget)

	if output == "" {
		t.Fatal("No html generated")
	}
	t.Log(output)
}
