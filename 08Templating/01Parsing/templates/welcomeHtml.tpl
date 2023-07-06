<!DOCTYPE HTML>
<HTML lang="en">

<head>
    <meta charset="UTF-8">
    <title>Hello World</title>
</head>
{{$Person:=.}}
<body>
    <h1>Fassak!</h1>
    <ul>
    {{range .}}
    <li> {{.}} </li> 
    {{end}}
    </ul>
</body>
</HTML>