package main

import (
	"bytes"
	"html/template"
	
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/html"

)

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

func HtmlMinify(content string) string {
    
	m := minify.New()
	m.AddFunc("text/html", html.Minify)

	s, err := m.String("text/html", content)
	if err != nil {
		panic(err)
	}

	return s
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
		<p> 
		<table>
		<tr><td><b>Alert Target Ids:</b></td><td></td></tr>
		{{range .TargetIds }}
			<tr>
				<td></td><td>{{ . }}</td>
			<tr>
		{{end}}
		</table>
		</p><br>
	{{ end }}
	{{ if .ImpactedServices }}
		<p></b>Impacted Services:</b>
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
	t, err := template.New("sh").Parse(tmpl)
	if err != nil {
		panic(err)
	}
	var buf bytes.Buffer
	err = t.Execute(&buf, serviceHealth)
	if err != nil {
		panic(err)
	}

	return HtmlMinify(buf.String())
}
