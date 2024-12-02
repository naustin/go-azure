package main

import ()

type Budget struct {
	Environment              string
	ThresholdType            string
	ForecastedTotalForPeriod string
	SpentAmount              string
	BudgetThreshold          string
	BudgetId                 string
	Description              string
	AzurePortalUrl           string
}

func CreateBudgetHtml(budgetValues Budget) string {
	tmpl := `
	<h1><b><strong>{{ .Environment }}: {{ .ThresholdType }} Budget Spend Amount Reached
	@{mul(div(if(equals('{{ .ThresholdType }}', 'Forecasted'),
	decimal(substring('{{ .ForecastedTotalForPeriod }}', 1)),
	decimal(substring('{{ .SpentAmount }}', 1))),
	decimal(substring('{{ .BudgetThreshold }}', 1))),
	100)}%</strong></b></h1><br>
<p>{{ .Description }}</p><br>
<p>Visit <a
href='https://{{ .AzurePortalUrl }}/#view/Microsoft_Azure_CostManagement/Menu/~/overview/open/budgets/openedBy/AzurePortal'>Cost
Management</a> for more details.</p><br>
<p><a href="https://{{ .AzurePortalUrl }}/#view/Microsoft_Azure_CostManagement/BudgetDetailBlade/budgetId/@{replace(substring('{{ .BudgetId }}', 1), '/' , '%2F' )}">Budget Link</a></p>
	`

	htmlText := BuildHtmlFromTemplate(tmpl, budgetValues)

	return HtmlMinify(htmlText)
}
