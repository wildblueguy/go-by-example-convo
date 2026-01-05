package main

import (
	"os"
	"text/template"
)

func main() {
	helloTemplate := template.New("Hello")
	helloTemplate, _ = helloTemplate.Parse("<h1>Hello {{.Name}}</h1>")
	helloTemplate.Execute(os.Stdout, map[string]string{"Name": "Gopher"})
}
