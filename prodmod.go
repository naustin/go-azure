package main

import ()

type ProdMod struct {
	Environment    string
	EventTimestamp string
	TargetIds      []string
}

func CreateProdModHtml(prodMod ProdMod) string {
	tmpl := `
	<h1><b><strong>{{ .Environment }}: Prod lock was removed</strong></b></h1><br>
	
	<p>At {{ .EventTimestamp }}, the resource group lock
    '{{ index .TargetIds 0 }}' was removed. If this is not intended (i.e the
    result of a production deployment), please consult the Cloud Infrastructure team.</p><br>
	
	`

	htmlText := BuildHtmlFromTemplate(tmpl, prodMod)

	return HtmlMinify(htmlText)
}
