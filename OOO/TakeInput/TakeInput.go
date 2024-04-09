package TakeInput

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func IpChar() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Confirm (y/n): ")
	r, _, err := reader.ReadRune()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("r: %v\n", string(r))
}

func IpString() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter Pattern: ")
	str, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln(err)
	}
	checker := strings.TrimSpace(str) == "007"
	if checker {
		fmt.Printf("Welcome Agent %v", str)
	} else {
		fmt.Println("No Entry")
	}
	fmt.Printf("Enter Name: ")
	str, err = reader.ReadString('\n')
	if err != nil {
		log.Fatalln(err)
	}
	if strings.TrimSpace(str[0:2]) == "Mr" {
		fmt.Println("Hello Sir!")
	} else {
		fmt.Println("Hello Madam!")
	}

}
