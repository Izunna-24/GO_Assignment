package main
import "fmt"

func main(){
	var (
		milesPerGallon1,
		milesPerGallon2,
		milesPerGallon3 float64
)
	for trips := 1; trips <= 3; trips++{
		if trips == 1{
			fmt.Println("Fisrt Trip")
		fmt.Print("How many miles was covered: ")
		var miles int
		fmt.Scanln(&miles)
		fmt.Print("How many galons was used: ")
		var galons int
		fmt.Scanln(&galons)
		
		 milesPerGallon1 = float64(miles) / float64(galons)
		 fmt.Println("Miles per gallon for trip 1:", milesPerGallon1)
	
	 }else if trips == 2{
		fmt.Println("Second Trip")
	fmt.Print("How many miles was covered: ")
		var miles int
		fmt.Scanln(&miles)
	fmt.Print("How many galons was used: ")
		var galons int
		fmt.Scanln(&galons)
		
		 milesPerGallon2 = float64(miles) / float64(galons)
		 fmt.Println("Miles per gallon for trip 2:", milesPerGallon2)

	}else if trips == 3{
		fmt.Println("third Trip")
	fmt.Print("How many miles was covered: ")
		var miles int
		fmt.Scanln(&miles)
	fmt.Print("How many galons was used: ")
		var galons int
		fmt.Scanln(&galons)
		
		 milesPerGallon3 = float64(miles) / float64(galons)
		 fmt.Println("Miles per gallon for trip 3:", milesPerGallon3)
 
	}
}
	 totalMilesPerGallon  := (milesPerGallon1 + milesPerGallon2 + milesPerGallon3)

	fmt.Println("Total miles Per gallon for all trips: ",totalMilesPerGallon,"miles/gallon") 

}