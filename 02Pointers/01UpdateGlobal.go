package main

import "fmt"

type itemAvailablity struct {
	globalState bool
}

type product struct {
	name      string
	code      uint8
	state     bool
	available *itemAvailablity
}

func (prdArry *product) changeStatus(change *itemAvailablity) {

	if change.globalState == false {
		prdArry.state = change.globalState
		fmt.Printf("Product Deactivated\n%v Not Available!\n", prdArry.name)
	} else {
		prdArry.state = change.globalState
		fmt.Printf("Product %v Available for Sale!\n", prdArry.name)
	}
}

func (prdArry *product) printProductDetails() {
	if prdArry.name != "" {
		fmt.Printf(
			"ProductName:\t%v\nProduct Code:\t%v\nActiveIncart\t%v\n",
			prdArry.name,
			prdArry.code,
			prdArry.state)
	}
}

func (prdArry *product) checkoutKart() {
	if prdArry.name != "" && prdArry.available.globalState == prdArry.state {
		fmt.Printf(
			"Purchased ProductName:\t%v\nProduct Code:\t%v\n",
			prdArry.name,
			prdArry.code)
	} else if prdArry.name != "" && prdArry.available.globalState != prdArry.state {
		fmt.Printf(
			"Not Available:\t%v\nProduct Code:\t%v\n",
			prdArry.name,
			prdArry.code)
	} else {
		println("")
	}
}

func printMultipleProducts(prdArry []product) {
	for i := 0; i < len(prdArry); i++ {
		var tempArr product = prdArry[i]
		tempArr.printProductDetails()
	}
}

func checkoutMultipleProducts(prdArry []product) {
	for i := 0; i < len(prdArry); i++ {
		var tempArr product = prdArry[i]
		tempArr.checkoutKart()
	}
}

func main() {

	var globalAv itemAvailablity
	globalAv.globalState = true

	var active itemAvailablity
	active.globalState = true

	var inactive itemAvailablity
	inactive.globalState = false

	shirtMet := make([]product, 5)
	shirtMet[0] = product{"Metallica", 101, active.globalState, &globalAv}
	shirtMet[1] = product{"Slayer", 111, active.globalState, &globalAv}
	shirtMet[2] = product{"Slipknot", 121, active.globalState, &globalAv}
	shirtMet[3] = product{"AmonAmarth", 131, active.globalState, &globalAv}
	shirtMet[3].changeStatus(&inactive)
	printMultipleProducts(shirtMet)
	checkoutMultipleProducts(shirtMet)
}
