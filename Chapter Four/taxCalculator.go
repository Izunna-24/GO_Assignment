ppackage main

import "fmt"


func main() {
	fmt.Print("Our system can determine deu tax of Three person's per time!!")
	for i := 0; i < 3; i++ {
		var name string
		var earning float64

		fmt.Print("Enter citizen's name: ")
		fmt.Scanln(&name)

		fmt.Printf("Enter %s's earning: ", name)
		fmt.Scanln(&earning)

		benchPay1 := 30000.0
		taxRate1 := 0.15
		taxRate2 := 0.20

		var tax float64

		if earning > benchPay1 {
			benchPayTax := earning * taxRate1
			benchPay2 := earning - benchPay1
			smallTax := benchPay2 * taxRate2
			tax = benchPayTax + smallTax
		} else {
			tax = earning * taxRate1
		}

		fmt.Printf("%s's Total Tax: %.2f\n", name, tax)
	}
}
