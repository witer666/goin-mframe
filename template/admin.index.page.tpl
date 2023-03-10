<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Gin</title>
    <style>
table
{
    width: 20%;
    border-collapse:collapse;
}
table,th, td
{
    border: 1px solid black;
}
    </style>
</head>
<body>
<h1>{{.title}}</h1>
<table>
{{range $k, $v := .list}}
<tr>
<td>
{{$v.uname}}
</td>
<td>
{{$v.sex}}
</td>
</tr>
{{end}}
</table>
</body>
</html>