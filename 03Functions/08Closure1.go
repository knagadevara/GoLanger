package main

import "fmt"

type mathER interface {
	compute()
}

type MathCompute func(a, b float32) float32

func computeR(toComp MathCompute, a float32, b float32) float32 {
	fmt.Printf("Things are getting added\n")
	return toComp(a, b)
}

func main() {

	var ComputeOpperands MathCompute
	ComputeOpperands = func(a, b float32) float32 {
		return a + b
	}
	fmt.Printf("Addition:\t%v\n", computeR(ComputeOpperands, 10, 20))
	fmt.Printf("Substraction:\t%v\n", computeR(func(a, b float32) float32 {
		return a - b
	}, 10, 20))

}
