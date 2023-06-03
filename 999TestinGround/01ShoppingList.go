package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	CentralTax float32 = 4.6
	StateTax   float32 = 6.5
)

type product struct {
	name     string
	price    float32
	quantity uint8
}

func calculateBill(itemsInCart []product) (uint8, float32) {

	var itemTotal, amountTotal float32
	var listedItems uint8

	println("Name\t\tPrice")
	println("--------\t--------")

	for i := 0; i < len(itemsInCart); i++ {
		var item product = itemsInCart[i]
		if item.name != "" {
			itemTotal = item.price * float32(item.quantity)
			fmt.Printf("%v\t%v\n", item.name, itemTotal)
			amountTotal += itemTotal
			listedItems += 1
		}
	}

	var taxAdditive float32 = amountTotal / (CentralTax + StateTax)
	println("--------\t--------")
	fmt.Printf("CST + SST \t%g + %g\nTotal Tax:\t%g\n", CentralTax, StateTax, taxAdditive)
	amountTotal = taxAdditive + amountTotal
	return listedItems, amountTotal

}

func getItemQuantity() int64 {
	fmt.Printf("Enter quantity of the List: ")
	inputReader := bufio.NewReader(os.Stdin)
	lstQty, err := inputReader.ReadString('\n')
	qty, err := strconv.ParseInt(strings.TrimSuffix(lstQty, "\n"), 0, 0)
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return 0
	}
	return qty
}

func main() {
	listOfProducts := make([]product, getItemQuantity())
	listOfProducts[0] = product{"Goldiee", 199, 5}
	listOfProducts[1] = product{"Nicotin", 99, 3}
	listedItems, amountTotal := calculateBill(listOfProducts)
	println("--------\t--------")
	fmt.Printf("TotalItems:\t%v\nTotalAmount:\t%v\n", listedItems, amountTotal)
}
