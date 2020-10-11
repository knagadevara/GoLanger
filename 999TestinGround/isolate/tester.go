package main

import "fmt"

func main() {
	buckets := make([][]string, 2)

	for i := 0; i <= 12; i++ {
		buckets = append(buckets, []string{})
	}

	fmt.Println(buckets)

}
