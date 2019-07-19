package main

import "fmt"

func main() {
	name := "Nikola Obradovic"

	tmpl := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Hello Go!</title>
	</head>
	<body>
	<h1>` + name + `</h1>
	</body>
	</html>
	`

	fmt.Println(tmpl)
}

// go run main.go > index.html
