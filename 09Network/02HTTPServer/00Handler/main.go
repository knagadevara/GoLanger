package main

import (
	"fmt"
	"log"
	"net/http"
)

type testType uint8

func (t *testType) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Default().Println("Running....")
	fmt.Fprintf(w, "Code to run")
}

func main() {
	var tst testType
	err := http.ListenAndServe(":8080", &tst)
	if err != nil {
		log.Fatalln(err)
	}
}
