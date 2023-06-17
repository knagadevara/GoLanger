package main

import "fmt"

// declaring an array

var MiNami [58]int

func main() {
	fmt.Printf("Binary\t\tUTF-8-No\tValue\n")
	fmt.Printf("------\t\t--------\t-----\n")
	for i := 65; i <= 122; i++ {
		MiNami[i-65] = i
		fmt.Printf("%b\t\t%d\t\t%c\n", MiNami[i-65], MiNami[i-65], MiNami[i-65])
	}
	fmt.Printf("Element of %T\n", MiNami[0])

}
