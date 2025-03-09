package main

import (
	"fmt"
)

func main() {
	var revenue float64
	var expenses float64
	var tax_rate float64
	
	fmt.Println("Enter the Revenue")
	fmt.Scan(&revenue)
	fmt.Println("Enter the Expenses")
	fmt.Scan(&expenses)
	fmt.Println("Enter the Tax Rate")
	fmt.Scan(&tax_rate)
	earningsBeforeTax := revenue - expenses
	earningsAfterTax := earningsBeforeTax * (1-(tax_rate/100))
	ratio := earningsBeforeTax / earningsAfterTax
	//formatting strings
	// fmt.Print("The profit is: ", earningsBeforeTax)
	// fmt.Print("The tax is: ", earningsAfterTax)
	// fmt.Print("The net profit is: ", ratio)
    fmt.Printf("The profit is: %.2f\n", earningsBeforeTax)
	fmt.Printf("The tax is: %.2f\n", earningsAfterTax)
	fmt.Printf("The net profit is: %.2f\n", ratio)
}

// output
// Enter the Revenue
// 10000
// Enter the Expenses
// 2500
// Enter the Tax Rate
// 30
// The profit is: 7500.00
// The tax is: 5250.00
// The net profit is: 1.43