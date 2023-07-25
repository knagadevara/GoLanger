package main

import (
	"fmt"
	"net/http"
	"uuid" // this was added by github.com/google/uuid
)

func createUUID(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("session")
	if err != nil {
		SessionUUID := uuid.NewString()
		cookie := &http.Cookie{Name: "session", Value: SessionUUID, HttpOnly: true}
		http.SetCookie(w, cookie)
	}
}

func rootPath(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context", "text/html;charset=utf-8")
	createUUID(w, r)
	fmt.Fprintf(w, "<html><head><h1>Created a Session</h1></head></html>")
}

func main() {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/favicon.ico", http.NotFound)
	serveMux.HandleFunc("/", rootPath)
	http.ListenAndServe(":8081", serveMux)
}
