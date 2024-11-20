package main

import ()

type ResourceHealth struct {
	Environment         string
	Title               string
	CurrentHealthStatus string
	EventTimestamp      string
	Level               string
	Issue               string
	Type                string
	Cause               string
	AzurePortalUrl      string
	TargetIds           []string
}

func CreateResourceHealthHtml(resourceHealth ResourceHealth) string {
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
	
		{{ if and .CurrentHealthStatus .TargetIds }}
		<p>At {{ .EventTimestamp }}, the resource '{{ index .TargetIds 0 }}' transitioned to a
				'{{ .CurrentHealthStatus }}' health state.
		</p>
		{{ end }}
	
		{{ if .Level }}
			<p><b>Level:</b> {{ .Level }}</p><br>
		{{ end }}
	
		{{ if .Issue }}
			<p><b>Issue:</b> {{ .Issue }}</p><br>
		{{ end }}
	
		{{ if .Type }}
		<p><b>Type:</b> {{ .Type }}</p><br>
		{{ end }}

		{{ if .Cause }}
		<p><b>Cause:</b> {{ .Cause }}</p><br>
		{{ end }}        

		<p><a href='https://{{ .AzurePortalUrl }}/resource{{ index .TargetIds 0 }}'>View Resource</a></p>

		</body></html>
	`

	htmlText := BuildHtmlFromTemplate(tmpl, resourceHealth)

	return HtmlMinify(htmlText)
}
