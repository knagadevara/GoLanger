package main

import (
	"log"
	"os"
)

func init() {
	nf, err := os.Create("log.txt")
	if err != nil {
		log.Fatalf("Error Cannot Create a log file: %v", err)
	}
	log.SetOutput(nf)
}

func main() {
	someFile, err := os.Open("NonExistingFile.txt")
	if err != nil {
		// log.Printf("Cannot Open file: %v\n", err)
		// log.Fatalf("Cannot Open file: %v\n", err)
		log.Panicf("Cannot Open file: %v\n", err)
	}
	defer someFile.Close()
}
