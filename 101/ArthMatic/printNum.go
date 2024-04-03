package ArthMatic

import "fmt"

func PrintNum() {
	for i := 40; i <= 126; i++ {
		fmt.Printf("decimal: %d \t binary: %b \t hexadecimal: %x \t utf-8: %q \n", i, i, i, i)
	}
}
