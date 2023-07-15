package main

import (
	"log"
	"net/http"
	"net/url"
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

func executeForm(w http.ResponseWriter, templateName string, method any) {
	err := parsedTemplate.ExecuteTemplate(w, templateName, method)
	if err != nil {
		log.Default().Fatalln(err)
	} else {
		log.Default().Println("Executed Template")
	}
}

func setRetrive(w http.ResponseWriter, r *http.Request, method string) {
	var formData url.Values
	w.Header().Set("Content-Type", "text/html ; charset=utf-8")
	executeForm(w, "parseValues.html", method)
	err := r.ParseForm()
	if err != nil {
		log.Default().Fatalln(err)
	} else {
		log.Default().Println("Executed ParsedForm, Populated values into PostForm and Form.")
	}
	switch method {
	case "GET":
		formData = r.Form
	case "POST":
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

func viaPost(w http.ResponseWriter, r *http.Request) {
	setRetrive(w, r, "POST")
}

func viaGet(w http.ResponseWriter, r *http.Request) {
	setRetrive(w, r, "GET")
}

func main() {
	servMux := http.NewServeMux()
	servMux.HandleFunc("/parseGet/", viaGet)
	servMux.HandleFunc("/parsePost/", viaPost)
	servMux.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8081", servMux)
}
