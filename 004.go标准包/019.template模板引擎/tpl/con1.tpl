{{/* 继承根模板 */}}
{{template "layout.tpl" .}}

{{/* 重新定义内容 */}}
{{define "content"}}
{{.content}}
{{end}}