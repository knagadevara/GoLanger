package main

import "fmt"

// declaring an array
// type aliases
// byte --> uin8
// rune --> int32

var MiNami [58]byte

func main() {
	var txR uint8
	fmt.Printf("Binary\t\tUTF-8-No\tValue\n")
	fmt.Printf("------\t\t--------\t-----\n")
	for txR = 65; txR <= 122; txR++ {
		MiNami[txR-65] = txR
		fmt.Printf("%b\t\t%d\t\t%c\n", MiNami[txR-65], MiNami[txR-65], MiNami[txR-65])
	}
	fmt.Printf("Element of %T\n", MiNami[0])

}
