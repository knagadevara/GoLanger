package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const somePara = "In this example, we created a slice, and then created a for loop that would iterate four times. Each iteration appended the current value of the loop variable i into the index of the numbers slice. However, this could lead to unnecessary memory allocations that could slow down your program. When adding to an empty slice, each time you make a call to append, the program checks the capacity of the slice. If the added element makes the slice exceed this capacity, the program will allocate additional memory to account for it. This creates additional overhead in your program and can result in a slower execution."

func main() {

	// Can print the line directly through
	// fmt.Println(somePara)
	// Converting the string into a readable reader Type
	reader := strings.NewReader(somePara)
	// Scanning over the reader object, placing them into a buffer and returning a Scanner Type
	scanner := bufio.NewScanner(reader)
	// split the entire string into words with delimiter.
	scanner.Split(bufio.ScanWords)
	// scanner.Scan() returns a boolean value telling that there is a string to scan
	fmt.Println(scanner.Scan())
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	err := scanner.Err()
	if err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

}
