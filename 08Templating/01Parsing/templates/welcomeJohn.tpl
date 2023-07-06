<!DOCTYPE HTML>
<HTML lang="en">

<head>
    <meta charset="UTF-8">
    <title>Hello World</title>
</head>
{{$x := .}}
<body>
    <h1>Fassak!</h1>
    <ul>
    <li> You've been painted by {{$x.FName}} , {{$x.LName}}  </li> 
    </ul>
</body>
</HTML>