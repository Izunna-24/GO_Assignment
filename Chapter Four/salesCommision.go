package main

import (
	"fmt"
)

func main() {
	var items int
	fmt.Println("To get the earnings enter 0 as items sold for last week\n")
	fmt.Print("Items sold for last Week: ")
	fmt.Scanln(&items)

	profit := 200.0
	var sumValues float64

	for items != 0 {
		var values float64
		fmt.Print("Value for items sold: ")
		fmt.Scanln(&values)
		sumValues += values
		fmt.Print("Items sold for last Week: ")
		fmt.Scanln(&items)
	}

	percentageProfit := 9.0 / 100

	earning := profit + (percentageProfit * sumValues)

	fmt.Printf("Salesperson's earning for last week: $%.2f\n", earning)
}
