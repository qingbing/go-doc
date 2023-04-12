<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{ .title }}</title>
    <style>
        *{
            margin: 0;
        }
        .nav{
            height: 50px;
            width: 100%;
            position: fixed;
            top: 0;
            background-color: beige;
        }
        .main{
            margin-top: 50px;
        }
        .menu{
            width: 20%;
            height: 100%;
            position: fixed;
            left: 0;
            background-color: antiquewhite;
        }
        .content{
            margin-left: 20%;
            width: 80%;
            height: 100%;
            position: fixed;
            background-color: azure;
        }
    </style>
</head>
<body>
    <div class="nav">nav</div>
    <div class="main">
        <div class="menu">menu</div>
        <div class="content">
            {{block "content" .}} {{end}}
        </div>
    </div>
</body>
</html>