package main

import (
	"html/template"
	"log"
	"net/http"
)

var parsedTemplate *template.Template

func init() {
	parsedTemplate = template.Must(template.ParseFiles("index.html.tpl"))
}

func executeGivenTemplate(w http.ResponseWriter, tempName, tempFlag string) {
	err := parsedTemplate.ExecuteTemplate(w, tempName, tempFlag)
	if err != nil {
		log.Default().Fatalln(err)
	}
}

func releaseOk(w http.ResponseWriter, r *http.Request) {
	executeGivenTemplate(w, "index.html.tpl", "Go-Ahead!")
	w.Header().Add("Release", "Ready")
	w.WriteHeader(200)
}

func releaseNo(w http.ResponseWriter, r *http.Request) {
	executeGivenTemplate(w, "index.html.tpl", "CalledOff!")
	w.Header().Add("Release", "Canceled")
	w.WriteHeader(404)
}

func releaseWait(w http.ResponseWriter, r *http.Request) {
	executeGivenTemplate(w, "index.html.tpl", "Awaiting...")
	w.Header().Add("Release", "Hold-Off")
}

func ver1() {
	var serveMux *http.ServeMux = http.NewServeMux() // creating a ServeMux ptr
	serveMux.HandleFunc("/", releaseWait)
	serveMux.HandleFunc("/wait/", releaseWait) // trailing "/" will handel /wait/for/this as well
	serveMux.HandleFunc("/ok", releaseOk)
	serveMux.HandleFunc("/no", releaseNo) // only handels /no and nothing else
	http.ListenAndServe(":8081", serveMux)
}

func ver2() {

	http.Handle("/", http.HandlerFunc(releaseWait))
	http.Handle("/wait/", http.HandlerFunc(releaseWait)) // trailing "/" will handel /wait/for/this as well
	http.Handle("/ok", http.HandlerFunc(releaseOk))
	http.Handle("/no", http.HandlerFunc(releaseNo)) // only handels /no and nothing else
	http.ListenAndServe(":8081", nil)
}

func main() {
	ver2()
}
