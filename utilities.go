package main

import (
	"bytes"
	"html/template"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/html"
)

func HtmlMinify(content string) string {

	m := minify.New()
	m.AddFunc("text/html", html.Minify)

	s, err := m.String("text/html", content)
	if err != nil {
		panic(err)
	}

	return s
}

func BuildHtmlFromTemplate(tmpl string, data any) string {

	t, err := template.New("html").Parse(tmpl)
	if err != nil {
		panic(err)
	}
	var buf bytes.Buffer
	err = t.Execute(&buf, data)
	if err != nil {
		panic(err)
	}

	return buf.String()
}
