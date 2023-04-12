<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>mutil</title>
</head>

<body>
    <h1>{{.title}}</h1>
    <hr>
    <h2>user struct</h2>
    <p>ID : {{.user.ID}}</p>
    <p>姓名 : {{.user.Name}}</p>
    <p>性别 : {{.user.Gender}}</p>
    <p>Call : {{.user.Say}}</p>
    <hr>
    <h2>map[string]any</h2>
    <p>ID : {{.mp.id}}</p>
    <p>姓名 : {{.mp.name}}</p>
    <p>性别 : {{.mp.gender}}</p>
    <p>index : {{index .mp "gender"}}</p>
    <hr>
    <h2>map with</h2>
    {{ with .mp}}
    <p>ID : {{.id}}</p>
    <p>姓名 : {{.name}}</p>
    <p>性别 : {{.gender}}</p>
    {{end}}
    <hr>
    <h2>map range</h2>
    {{range $key,$value := .mp}}
    <p>key: {{$key}}, value: {{$value}}</p>
    {{else}}
    <p>空数据</p>
    {{end}}
    <hr>
    <p>myfunc1: {{myfunc1 "world"}}</p>
    <p>myfunc2: {{myfunc1 "xxx"}}</p>
</body>

</html>