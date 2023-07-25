package main

import (
	"fmt"
	"net/http"
	"uuid"
)

func createUUID(w http.ResponseWriter, r *http.Request) {
	SessionUUID := uuid.NewString()
	http.SetCookie(w, &http.Cookie{Name: "session", Value: SessionUUID, HttpOnly: true})
}

func rootPath(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context", "text/html;charset=utf-8")
	createUUID(w, r)
	fmt.Fprintf(w, "<html><head><h1>Created a Session</h1></head></html>")
}

func main() {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/", rootPath)
	http.ListenAndServe(":8081", serveMux)
}
