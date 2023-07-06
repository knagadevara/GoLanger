<!DOCTYPE HTML>
<HTML lang="en">

<head>
    <meta charset="UTF-8">
    <title>Hello World</title>
</head>
<body>
    <ul>
    <h2>Location -> Killed</h2>
        {{range $k , $v := .}}
        <li>{{$k}} -> {{$v}}</li>
        {{end}}
    </ul>

</body>
</HTML>