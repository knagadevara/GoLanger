package main

import "fmt"

type mathER interface {
	compute()
}

type MathCompute func(a, b float32)            // can be used as a callback function
type MathComputeRet func(a, b float32) float32 //returns a float32 hence can be used as a closure

func computeWrappeR(toComp MathComputeRet, a, b float32) float32 {
	fmt.Printf("Compute\t")
	return toComp(a, b)
}

func printDetails(funcName string, cllbackFunc MathCompute, a, b float32) {
	fmt.Printf("%s:\t", funcName)
	cllbackFunc(a, b)
}

func main() {

	// making a closure which returns the memory address of type function intrun returning float32
	// example on function expression
	var add2NumR MathComputeRet
	add2NumR = func(a, b float32) float32 {
		return a + b
	}
	var sub2NumR MathComputeRet
	sub2NumR = func(a, b float32) float32 {
		return a - b
	}
	var mul2NumR MathComputeRet
	mul2NumR = func(a, b float32) float32 {
		return a * b
	}

	// Making a callback function wich directly executes but not returns anything
	var add2Num MathCompute
	add2Num = func(a, b float32) {
		fmt.Printf("%.2f\n", a+b)
	}
	var sub2Num MathCompute
	sub2Num = func(a, b float32) {
		fmt.Printf("%.2f\n", a-b)
	}
	var mul2Num MathCompute
	mul2Num = func(a, b float32) {
		fmt.Printf("%.2f\n", a*b)
	}
	fmt.Printf("The below functions are closures run as function expressions.\n")
	fmt.Printf("Addition:\t%.2f\n", computeWrappeR(add2NumR, 10, 20))
	fmt.Printf("Substraction:\t%.2f\n", computeWrappeR(sub2NumR, 10, 20))
	fmt.Printf("Multiplication:\t%.2f\n", computeWrappeR(mul2NumR, 10, 20))
	fmt.Printf("Division:\t%.2f\n", computeWrappeR(func(a, b float32) float32 { return a / b }, 10, 20))

	fmt.Printf("The below functions are callbacks run as function expressions.\n")
	printDetails("Addition", add2Num, 1, 2)
	printDetails("Substraction", sub2Num, 1, 2)
	printDetails("Multiplication", mul2Num, 1, 2)
}
