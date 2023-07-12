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
var err error

func init() {
	parsedTemplate = template.Must(template.ParseFiles("index.html"))
}

func (p *authPerson) ServeHTTP(httpRespWrtr http.ResponseWriter, httpReq *http.Request) {
	// Parse the request
	// ParseForm populates httpReq.Form and httpReq.PostForm.
	err = httpReq.ParseForm()
	if err != nil {
		log.Default().Fatalln(err)
	}

	err = parsedTemplate.ExecuteTemplate(httpRespWrtr, "index.html", httpReq.Form)
	if err != nil {
		log.Default().Fatalln(err)
	}

}

func main() {
	var jw authPerson
	http.ListenAndServe(":8081", &jw)
}
