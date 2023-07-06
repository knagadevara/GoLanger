<!DOCTYPE HTML>
<HTML lang="en">

<head>
    <meta charset="UTF-8">
    <title>Hello World</title>
</head>
<body>
    <h1>Fassak!</h1>
<ul>
{{range .}}
{{$Car:=.Car}}{{$Owner:=.Owner}} <li> A  {{$Car.Doors}} Doors {{$Car.Name}} was made especially by {{$Car.Manufacturer}} to honor {{$Owner.LName}} , {{$Owner.FName}} for his {{$Owner.Age}} Birthday</li>
{{end}}
</ul>
</body>
</HTML>