package Models

const WikiPageTemplate = `

Dear {{.Name}},

{{if .Attended}}

It was a pleasure to see you at the wedding.

{{- else}}

It is a shame you couldn't make it to the wedding.

{{- end}}

{{with .Gift -}}

Thank you for the lovely {{.}}.

{{end}}

Best wishes,

Josie

`
