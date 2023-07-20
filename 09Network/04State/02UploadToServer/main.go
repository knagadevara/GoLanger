package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var parsedTemplate *template.Template

func checkError(w http.ResponseWriter, err error) {
	if w != nil {
		if err != nil {
			log.Default().Fatalln(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		if err != nil {
			log.Default().Fatalln(err)
		}
	}
}

func init() {
	parsedTemplate = template.Must(template.ParseGlob("*.gohtml"))
}

func executeForm(w http.ResponseWriter, templateName string, data any) {
	err := parsedTemplate.ExecuteTemplate(w, templateName, data)
	checkError(w, err)
}

func readFile(w http.ResponseWriter, r *http.Request) {
	var str string
	if r.Method == http.MethodPost {
		multiPartFile, fileHeader, err := r.FormFile("fileName")
		checkError(w, err)
		defer multiPartFile.Close()
		bs, err := io.ReadAll(multiPartFile)
		checkError(w, err)
		newFile, err := os.Create(filepath.Join("./store/", fileHeader.Filename))
		checkError(w, err)
		defer newFile.Close()
		_, err = newFile.Write(bs)
		checkError(w, err)
		str = string(bs)
	}
	w.Header().Set("Content-Type", "text/html ; charset=utf-8")
	w.WriteHeader(http.StatusAccepted)
	executeForm(w, "index.gohtml", str)
}

func main() {
	servMux := http.NewServeMux()
	servMux.HandleFunc("/", readFile)
	servMux.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(":8081", servMux)
	checkError(nil, err)
}
