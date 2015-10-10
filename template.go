package main

import (
	"html/template"
)

const (
	baseHTML string = `
		{{define "html"}}
			<html>
				<head>
					<meta charset="utf-8">
					<title>{{template "title" .}}</title>
					<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/normalize/3.0.3/normalize.min.css">
				</head>
				<body>
					<div class="container">
						{{template "body" .}}
					</div>
				</body>
			</html>
		{{end}}
	`
	directoryHTML string = `
		{{define "title"}}{{.Title}}{{end}}
		{{define "body"}}
			<h1>{{.Title}}</h1>
			<hr>
			<ul>
				{{ range $item := .Items }}
					<li><a href="{{ $item }}">{{ $item }}</a></li>
				{{ end }}
			</ul>
			<hr>
		{{end}}
	`
	notFoundHTML string = `
		{{define "title"}}Error response{{end}}
		{{define "body"}}
			<h1>Error response</h1>
			<p>Error code 404.</p>
			<p>Message: File not found.</p>
			<p>Error code explanation: 404 = Nothing matches the given URI.</p>
		{{end}}
	`
)

var directoryTemplate *template.Template
var notFoundTemplate *template.Template

func init() {
	directoryTemplate, _ = template.New("directory").Parse(baseHTML + directoryHTML)
	notFoundTemplate, _ = template.New("notFound").Parse(baseHTML + notFoundHTML)
}
