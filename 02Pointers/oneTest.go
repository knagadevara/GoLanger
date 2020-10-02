package main

import "fmt"

var (
	a int  = 10
	b      = a
	c *int = &a
	d      = &b
	e *int = &a
	f int
)

func main() {

	fmt.Printf("%v , %T \n", a, a) // 10
	fmt.Printf("%v , %T \n", b, b) // 10
	fmt.Printf("%v , %T \n", c, c) // memory address of a
	fmt.Printf("%v , %T \n", d, d) // memory address of b but will it be as same as a?
	fmt.Printf("%v , %T \n", e, e) // ??
	f = *c + *d
	fmt.Printf("%v , %T \n", f, f) // 10

}
