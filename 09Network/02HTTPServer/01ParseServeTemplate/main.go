package main

import (
	"html/template"
	"log"
	"net/http"
)

type authPerson struct {
	FName string
	LName string
	Age   uint8
}

var parsedTemplate *template.Template

func (p *authPerson) ServeHTTP(httpRespWrtr http.ResponseWriter, ptrHttpReq *http.Request) {
	tmpErr := parsedTemplate.ExecuteTemplate(httpRespWrtr, "index.html.tpl", p)
	if tmpErr != nil {
		log.Default().Fatalln(tmpErr)
	}
	log.Default().Println("Parsed Successfully")
}

func init() {
	parsedTemplate = template.Must(template.ParseFiles("index.html.tpl"))
}

func main() {
	var jW authPerson
	jW.Age = 32
	jW.FName = "John"
	jW.LName = "Wick"
	http.ListenAndServe(":8081", &jW)
}

// Substitutes the data into template and serves it to connections
