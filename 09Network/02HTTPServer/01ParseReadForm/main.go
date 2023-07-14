package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strconv"
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

func (p *authPerson) printDetails() {
	fmt.Printf("Name:%s\t\nSurName:%s\t\nAge:%d\t\n", p.FName, p.LName, p.Age)
}

func (p *authPerson) mapFormValues(reqForm *url.Values) {
	Age, err := strconv.ParseUint(reqForm.Get("Age"), 10, 8)
	if err != nil {
		log.Default().Fatalln(err)
	}
	p.Age = uint8(Age)
	p.FName = reqForm.Get("FirstName")
	p.LName = reqForm.Get("LastName")
	p.printDetails()
}

func (p *authPerson) setHeadderValues(hdr *http.Header) {
	hdr.Set("API-Version", "3")
	hdr.Set("API-KEY", "ljdbfsliflsnwflsdnf")
	hdr.Set("UserAuthValidity", "24h")
	hdr.Set("user", p.FName)
}

func (p *authPerson) ServeHTTP(httpRespWrtr http.ResponseWriter, httpReq *http.Request) {
	// request will be modified in the coming step.
	// ParseForm function will Parse the data from body/url and
	// populates it to httpReq.Form and httpReq.PostForm which are of typr url.Values (map[string][]string).
	err = httpReq.ParseForm()
	if err != nil {
		log.Default().Fatalln(err)
	}
	err = parsedTemplate.ExecuteTemplate(httpRespWrtr, "index.html", httpReq)
	if err != nil {
		log.Default().Fatalln(err)
	}
	p.mapFormValues(&httpReq.PostForm)
	hddr := httpRespWrtr.Header()
	p.setHeadderValues(&hddr)
	log.Default().Println("\tConnected...")
}

func main() {
	var jw authPerson
	http.ListenAndServe(":8081", &jw)
}
