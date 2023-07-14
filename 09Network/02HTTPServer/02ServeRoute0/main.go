package main

import (
	"fmt"
	"net/http"
)

type person int

func (p *person) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/OK":
		fmt.Fprintf(w, "Go-Ahead!")
		w.Header().Add("Release", "Ready")
		w.WriteHeader(200)
	case "/NF":
		fmt.Fprintf(w, "CalledOff!")
		w.Header().Add("Release", "Canceled")
		w.WriteHeader(404)
	default:
		w.Header().Add("Release", "Hold-Off")
		fmt.Fprintf(w, "Awaiting...")
	}
}

func main() {
	var p person
	http.ListenAndServe(":8081", &p)
}
