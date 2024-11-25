package main

import ()

type Metric struct {
	Environment      string
	AlertRule        string
	MonitorCondition string
	EventTimestamp   string
	Description      string
	AzurePortalUrl   string
	TargetIds        []string
}

func CreateMetricHtml(metric Metric) string {
	tmpl := `
<h1><b><strong>{{ .Environment }}: Metric Alert was {{ .MonitorCondition }}</strong></b></h1>
<br>
<p><b>Alert rule:</b> '{{ .AlertRule }}'</p><br>
<p><b>Resource:</b> '{{ index .TargetIds 0 }}'</p><br>
<p><b>Event Time:</b> {{ .EventTimestamp }}</p><br>
<p><b>Alert rule description:</b> {{ .Description }}</p><br>
<p><b>Resource ID:</b> {{ index .TargetIds 0 }}<br>
<a href='https://{{ .AzurePortalUrl }}/resource{{ index .TargetIds 0 }}'>View Resource ></a></p><br>
<p>	
	`

	htmlText := BuildHtmlFromTemplate(tmpl, metric)

	return HtmlMinify(htmlText)
}
