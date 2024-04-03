package ArthMatic

import (
	"fmt"
	"math"
)

var (
	Hour        byte // alias to uint8
	Age         uint8
	Phone       uint32
	Salary      float32
	Temperature int16
)

func premetiveNumbers() {
	Hour = 12
	Age = 10
	Phone = 123456789
	Salary = 123456.789
	Temperature = -375

	fmt.Println("Hour: ", Hour)
	fmt.Println("Age: ", Age)
	fmt.Println("Phone: ", Phone)
	fmt.Println("Salary: ", Salary)
	fmt.Println("Temperature: ", Temperature)

}

func BasicArth() {
	val1 := 11 // automatic type would be assigned as int
	val2 := 31 // int is alias to int 64

	fmt.Printf("The original valuse inside Val1: %v, Val2: %v will not change by below Ops\n", val1, val2)
	fmt.Printf("Addition: %v\n", val2+val1)
	fmt.Printf("Substraction: %v\n", val2-val1)
	fmt.Printf("Multiplication: %v\n", val2*val1)
	fmt.Printf("Division: %v\n", val2/val1)
	fmt.Printf("Modulus: %v\n", val2%val1) // Modulus not defined for Float
	fmt.Printf("The original valuse inside Val1: %v, Val2: %v will change by Incrementing/Decrementing Ops\n", val1, val2)
	fmt.Printf("Incrementing val1++ : %v\n", val1)
	fmt.Printf("Decrementing val2-- : %v\n", val1)

	val3 := 122.34
	fmt.Println("Square Root", math.Sqrt(val3))
}
