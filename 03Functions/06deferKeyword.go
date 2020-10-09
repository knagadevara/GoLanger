package main

import "fmt"

// Hello Demo
// func hello() string {
func hello() {
	// return fmt.Sprintf("Hello")
	fmt.Println("Hello")
}

// World Demo
// func world() string {
func world() {
	// return fmt.Sprintf("World")
	fmt.Println("World")
}

// ExecuMe differs the execution of world till the function reaches its end.
func ExecuMe() {
	defer world()
	hello()
}
