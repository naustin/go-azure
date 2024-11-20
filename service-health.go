package main

import ()

type ServiceHealth struct {
	Environment          string
	Title                string
	Communication        string
	IncidentType         string
	ImpactStartTime      string
	ImpactMitigationTime string
	TargetIds            []string
	ImpactedServices     []string
}

func CreateServiceHealthHtml(serviceHealth ServiceHealth) string {
	tmpl := `
	<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" 
          content="width=device-width, initial-scale=1.0">
    <title>Table Cell min-width</title>
    <style>
        td {
            min-width: 100px;
            padding: 8px;
        }
    </style>
</head>
</body>
	<h1><b><strong>{{ .Environment }}: {{ .Title }}</strong></b></h1><br>
	<p>{{ .Communication }}</p><br>
	{{ if .IncidentType }}
		<p><b>Type:</b> {{ .IncidentType }}</p><br>
	{{ end }}
	{{ if .ImpactStartTime }}
		<p><b>Start Time:</b> {{ .ImpactStartTime }}</p><br>
	{{ end }}
	{{ if .ImpactMitigationTime }}
		<p><b>Mitigation Time:</b> {{ .ImpactMitigationTime }}</p><br>
	{{ end }}
	{{ if .TargetIds }}
		<p><b>Alert Target Ids:</b>
		<table>
		{{range .TargetIds }}
			<tr>
				<td>{{ . }}</td>
			<tr>
		{{end}}
		</table>
		</p>
	{{ end }}
	{{ if .ImpactedServices }}
		<p><b>Impacted Services:</b>
		<table>
		{{range .ImpactedServices }}
			<tr>
				<td>{{ . }}</td>
			<tr>
		{{end}}	
		</table>
		</p>
	{{ end }}
	</body></html>
	`

	htmlText := BuildHtmlFromTemplate(tmpl, serviceHealth)

	return HtmlMinify(htmlText)
}
