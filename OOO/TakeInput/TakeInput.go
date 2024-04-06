package TakeInput

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func GiveInput() {
	reader := bufio.NewReader(os.Stdin)
	r, _, err := reader.ReadRune()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("r: %v\n", r)
}
