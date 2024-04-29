package main

import "fmt"

func main() {
	
	for count := 1; count <= 10; count++ {
		for j := 1; j <= count; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}


	for count := 10; count >= 1; count-- {
		for j := 1; j <= count; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
