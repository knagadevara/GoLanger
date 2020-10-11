package main

import "fmt"

// declaring an array

var MiNami [70]string

func main() {

	for i := 65; i <= 134; i++ {

		MiNami[i-65] = string(i)
	}

	for _, value := range MiNami {

		fmt.Printf("Binary: %b \n", value)
	}
}
