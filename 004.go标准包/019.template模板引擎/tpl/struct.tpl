<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Struct</title>
</head>
<body>
    {{/* 注释，可以换行*/}}
    <p>ID : {{.ID}}</p>
    <p>姓名 : {{.Name}}</p>
    <p>性别 : {{.Gender}}</p>
    <p>Len(性别) : {{len .Gender}}</p>
    <p>Call : {{.Say}}</p>
    
    <hr>
    {{ $id:=.ID }}
    <p>custom id:    {{- $id -}}</p>
    {{ $str:="test name" }}
    <p>custom var: {{ $str }}</p>
</body>
</html>