package main

import (
	"fmt"
	"os"
)

func OpenReadFile(fileName string) {
	fileIO, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("%v", err)
		// return "", err
	}
	defer fileIO.Close()
	readBuffer := make([]byte, 0, 30)
	readBytes, err := fileIO.Read(readBuffer)
	if err != nil {
		fmt.Printf("%v", err)
		// return "", err
	}
	fmt.Printf("%v\n", readBytes)
	// return readBytes, err

}

func main() {
	fmt.Printf("Starts to Execute\n")
	OpenReadFile("exampleFile")
}
