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
    EBT := fmt.Sprintf("The profit is: %.2f\n", earningsBeforeTax)
	EAT := fmt.Sprintf("The tax is: %.2f\n", earningsAfterTax)
	ratio_str := fmt.Sprintf("The net profit is: %.2f\n", ratio)
	fmt.Print(EBT,EAT,ratio_str)
}