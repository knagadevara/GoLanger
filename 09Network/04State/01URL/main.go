package main

import (
	"log"
	"net/http"
	"net/url"
	"strings"
	"text/template"
)

type productUser struct {
	FirstName string
	LastName  string
	Age       string
}

var parsedTemplate *template.Template

func init() {
	parsedTemplate = template.Must(template.ParseGlob("*.html"))
}

func executeForm(w http.ResponseWriter, templateName string, data any) {
	err := parsedTemplate.ExecuteTemplate(w, templateName, data)
	if err != nil {
		log.Default().Fatalln(err)
	} else {
		log.Default().Println("Executed Template")
	}
}

func setRetrive(w http.ResponseWriter, r *http.Request) {
	var formData url.Values
	method := strings.Split(string(r.URL.Path), "/")[1]
	log.Default().Println(method)
	w.Header().Set("Content-Type", "text/html ; charset=utf-8")
	executeForm(w, "parseValues.html", method)
	err := r.ParseForm()
	if err != nil {
		log.Default().Fatalln(err)
	} else {
		log.Default().Println("Executed ParsedForm, Populated values into PostForm and Form.")
	}
	switch r.Method {
	case http.MethodGet:
		formData = r.Form
	case http.MethodPost:
		formData = r.PostForm
	default:
		formData.Add("FirstName", r.FormValue("FirstName"))
		formData.Add("LastName", r.FormValue("LastName"))
		formData.Add("Age", r.FormValue("Age"))
	}
	var endUser productUser = productUser{
		formData.Get("FirstName"),
		formData.Get("LastName"),
		formData.Get("Age"),
	}
	executeForm(w, "redirectData.html", endUser)
}

func main() {
	servMux := http.NewServeMux()
	servMux.HandleFunc("/get", setRetrive)
	servMux.HandleFunc("/post", setRetrive)
	servMux.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8081", servMux)
}
