package main

import "fmt"

func PrintMyName(name string) func() string {
	return func() string {
		return fmt.Sprintf("Hello %s", name)
	}
}

func main() {

	var getName func() string = PrintMyName("SivaRam")
	
	fmt.Printf("Type: %T \n", getName)
	fmt.Printf("Type: %T \n", getName())
	fmt.Printf("Value %v \n", getName())

}
