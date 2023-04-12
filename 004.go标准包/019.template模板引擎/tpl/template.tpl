<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>嵌套模板</title>
</head>
<body>
    <h1>测试嵌套template语法</h1>
    <hr>
    {{/* 嵌套单独的模板文件 */}}
    {{ template "ul.tpl" }}
    <hr>
    {{/* 嵌套 define 定义的模板 */}}
    {{ template "ol.tpl" }}
</body>
</html>

{{ define "ol.tpl"}}
<ol>
    <li>吃饭</li>
    <li>睡觉</li>
    <li>打豆豆</li>
</ol>
{{end}}