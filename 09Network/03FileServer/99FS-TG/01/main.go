package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var parsedTemplate *template.Template

func init() {
	parsedTemplate = template.Must(template.ParseFiles("dog.html"))
}

func execTemplate(w http.ResponseWriter, tempName, tempFlag string) {
	err := parsedTemplate.ExecuteTemplate(w, tempName, tempFlag)
	if err != nil {
		http.Error(w, "Content Not found!", http.StatusNotFound)
		log.Default().Fatalln(err)
	}
}

func foo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html ; charset=utf-8")
	execTemplate(w, "dog.html", "In The ROOT!")
}

func serveDog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html ; charset=utf-8")
	execTemplate(w, "dog.html", "Woof Wof Wof..")
}

func revealDog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html ; charset=utf-8")
	fmt.Fprintln(w, `<img src="/pics/dog.jpg" width="500" height="600" style="vertical-align:middle">`)
	// http.ServeFile(w, r, "dog.jpg")
}

func main() {
	servMux := http.NewServeMux()
	rootFileDir := http.Dir("/opt/GoLearner/GoLanger/09Network/03FileServer/CommonObjects/images/")
	fileServer := http.FileServer(rootFileDir)
	mountedStrippedPrifixHandler := http.StripPrefix("/pics/", fileServer)
	servMux.HandleFunc("/", foo)
	// http.Handle("/pics", srvRtDrHndlr) // can only serve on "/" root and the implementation will not work, hence we need to strip the pattern
	servMux.Handle("/pics/", mountedStrippedPrifixHandler)
	servMux.HandleFunc("/dog/", serveDog)
	servMux.HandleFunc("/who/", revealDog)
	servMux.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8081", servMux)
}
