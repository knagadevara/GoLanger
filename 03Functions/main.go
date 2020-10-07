package main

import "fmt"

func main() {

	fmt.Println("Generic Typed returns")
	fmt.Println(StudentDetails("Karthik", 2013, 21100024))
	fmt.Println("Named Returns")
	fmt.Println(Minami("karthik", 29))
	var numList = []float64{32, 141, 154, 11, 54, 111, 1788}
	// avg := MyAverage(numList)
	// Does not work because it is only allowed to take a 'float64' type but not a slice/list of 'float64'.
	// the correct way to pass it as an argument is by unpacking the variable
	// syntax:
	// var variableX = []string{ "Karthik" , "Sai" , "Venkata" , "Dev" }
	// function(variableX...)
	avg2 := MyAverage(numList...)
	avg := MyAverage(32, 141, 154, 11, 54, 111, 1788)
	fmt.Printf("The average value of the directly passed arguments MyAverage(32, 141, 154, 11, 54, 111, 1788) is %v\n", avg)
	fmt.Printf("The average value of the slice which is used by un-packing the slice while passing MyAverage(numList...) is %v\n", avg2)

	WhatsMyName := SayMyName()
	fmt.Printf("%T \n", WhatsMyName())
	fmt.Printf("%v \n", WhatsMyName())

	WhatsYourAge := WhatsYourAge()
	fmt.Printf("%T \n", WhatsYourAge)
	fmt.Printf("%v \n", WhatsYourAge)
	fmt.Printf("%p \n", &WhatsYourAge)
	fmt.Printf("%v \n", WhatsYourAge())
}
