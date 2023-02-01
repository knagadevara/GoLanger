package main

import (
	"fmt"
)

const (
	NoTaxZone  float64 = 700000.0
	slab_0_ul  float64 = 300000.0
	slab_0_tax float64 = 0.0

	slab_1_ul  float64 = 600000.0
	slab_1_tax float64 = 15000.0

	slab_2_ul  float64 = 900000.0
	slab_2_tax float64 = 30000.0

	slab_3_ul  float64 = 1200000.0
	slab_3_tax float64 = 45000.0

	slab_4_ul  float64 = 1500000.0
	slab_4_tax float64 = 90000.0
)

func TaxCalculator(salary float64) (TDS float64) {

	var totalTax float64
	fmt.Println(`
	TaxCalculator!
	Please find the below list.
	0 - 300000 = 0
	300000  - 600000    05%   = 15,000
	600000  - 900000    10%   = 30,000
	900000  - 1200000   15%   = 45,000
	1200000 - 1500000   20%   = 90,000	
	`)

	if salary <= NoTaxZone {
		fmt.Printf("Congratulations: You Pay No tax!")
		totalTax = slab_0_tax
	} else {
		if salary > NoTaxZone && salary <= slab_2_ul {
			totalTax = ((salary - slab_0_ul) * 0.05) + ((salary - slab_1_ul) * 0.10)

		} else if salary > slab_2_ul && salary <= slab_3_ul {
			totalTax = ((salary - slab_0_ul) * 0.05) + ((salary - slab_1_ul) * 0.10) + ((salary - slab_2_ul) * 0.15)

		} else if salary > slab_3_ul && salary <= slab_4_ul {
			totalTax = ((salary - slab_0_ul) * 0.05) + ((salary - slab_1_ul) * 0.10) + ((salary - slab_2_ul) * 0.15) + ((salary - slab_3_ul) * 0.20)
		} else if salary >= slab_4_ul {
			totalTax = ((salary - slab_0_ul) * 0.05) + ((salary - slab_1_ul) * 0.10) + ((salary - slab_2_ul) * 0.15) + ((salary - slab_3_ul) * 0.20) + ((salary - slab_4_ul) * 0.30)
		}
	}
	return totalTax
}

func main() {

	var amount float64
	fmt.Println("Please enter your Salary: ")
	fmt.Scanf("%f", &amount)
	fmt.Printf("\nTDS on your Salary: %f\n", TaxCalculator(amount))

}
