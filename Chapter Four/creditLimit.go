package main
import "fmt"

func main(){
	fmt.Print("Enter account number: ")
	var number string
	fmt.Scanln(&number)

	fmt.Print("balance at the beginning of the month: ")
	var balance int
	fmt.Scanln(&balance)

	fmt.Print("total of items charged by customer this month: ")
	var items int
	fmt.Scanln(&items)

	fmt.Print("total credit applied to customer this month: ")
	var credit int
	fmt.Scanln(&credit)

	fmt.Print("credit limit: ")
	var limit int
	fmt.Scanln(&limit)

	newBalance := balance + items - credit

		if newBalance <= credit {
	fmt.Println("User with account number", number, " has a new balance of ", newBalance)
		}else{
			fmt.Println("credit limit exceeded")
		}
}