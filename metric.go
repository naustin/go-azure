package main

import ()

type Metric struct {
	Environment         string
	AlertRule      string
	EventTimestamp string
	Description string
	AzurePortalUrl string
	TargetIds           []string
}

func CreateMetricHtml(metric Metric) string {
	tmpl := `
<h1><b><strong>Metric Alert was @{triggerBody()?['data']?['essentials']?['monitorCondition']}</strong></b></h1>
<br>
<p><b>Alert rule:</b> '{{ .AlertRule }}'</p><br>
<p><b>Resource:</b> '{{ index .TargetIds 0 }}'</p><br>
<p><b>Event Time:</b>{{ .EventTimestamp }}</p><br>
    
<p><b>Alert rule description:</b> {{ .Description }}</p><br>
<p><b>Resource ID:</b> {{ {{ index .TargetIds 0 }} }}<br>
<a href='https://{{ .AzurePortalUrl }}/resource{{ index .TargetIds 0 }}'>View Resource ></a></p><br>
<p>	
<h2>Alert Activated Because</h2>
</p><br>
<p>Metric Name: @{if(equals(triggerBody()?['data']?['essentials']?['signalType'], 'Log'),
    triggerBody()?['data']?['alertContext']?['condition']?['allOf'][0]?['metricMeasureColumn'],
    triggerBody()?['data']?['alertContext']?['condition']?['allOf'][0]?['metricName'])}</p><br>
<p>Metric Namespace: @{if(equals(triggerBody()?['data']?['essentials']?['signalType'], 'Log'),
    triggerBody()?['data']?['alertContext']?['condition']?['allOf'][0]?['targetResourceTypes'],
    triggerBody()?['data']?['alertContext']?['condition']?['allOf'][0]?['metricNamespace'])}</p><br>
<p>Time Aggregation: @{triggerBody()?['data']?['alertContext']?['condition']?['allOf'][0]?['timeAggregation']}</p><br>
<p>Period: @{triggerBody()?['data']?['alertContext']?['condition']?['windowSize']}</p><br>
<p>Value: @{triggerBody()?['data']?['alertContext']?['condition']?['allOf'][0]?['metricValue']}</p><br>
<p>Operator: @{triggerBody()?['data']?['alertContext']?['condition']?['allOf'][0]?['operator']}</p><br>
<p>Threshold: @{triggerBody()?['data']?['alertContext']?['condition']?['allOf'][0]?['threshold']}</p><br>
<p>@{if(equals(triggerBody()?['data']?['essentials']?['signalType'], 'Log'), concat('<a href=''', triggerBody()?['
        data']?['alertContext']?['condition']?['allOf'][0]?['linkToSearchResultsUI'], '''>Search Results</a>' ), ''
        )}</p>
	`

	htmlText := BuildHtmlFromTemplate(tmpl, metric)

	return HtmlMinify(htmlText)
}