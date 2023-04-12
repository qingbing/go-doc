<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Xss</title>
    
</head>
<body>
    <h1>脚本</h1>
    <p>{{ .script }}</p>
    <h2 id="test"></h2>
    {{ safe .safeScript}}
    <h2 id="test2"></h2>
    {{ .safeScript2 | safe}}
</body>
</html>