package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 3; i++ {
		var name string
		var earning float64

		fmt.Print("Enter citizen's name: ")
		fmt.Scanln(&name)

		fmt.Printf("Enter %s's earning: ", name)
		fmt.Scanln(&earning)

		tax := calculateTax(earning)
		fmt.Printf("%s's Total Tax: %.2f\n", name, tax)
	}
}

func calculateTax(earning float64) float64 {
	benchPay1 := 30000.0
	taxRate1 := 0.15
	taxRate2 := 0.20

	if earning > benchPay1 {
		benchPayTax := earning * taxRate1
		benchPay2 := earning - benchPay1
		smallTax := benchPay2 * taxRate2
		return benchPayTax + smallTax
	}

	return earning * taxRate1
}
