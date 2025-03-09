package main

import "fmt"
func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	var r = getRevenue(revenue)
	var e = getExpenses(expenses)
	var t = getTaxRate(taxRate)
    
	// var ebt,profit,ratio = calculate(r,e,t)
    
	ebt,profit,ratio := calculate(r,e,t)
	fmt.Print("EBT: ", ebt, "Profit: ", profit, "Ratio: ", ratio)

}
// functions to accept user inputs and calculate 
func getRevenue (revenue float64) float64 {
	fmt.Println("Enter the revenue:")
	fmt.Scan(&revenue)
	return revenue
}

func getExpenses (expenses float64) float64 {
	fmt.Println("Enter the expenses:")
	fmt.Scan(&expenses)
	return expenses
}

func getTaxRate (taxRate float64) float64 {
	fmt.Println("Enter the taxRate:")
	fmt.Scan(&taxRate)
	return taxRate
}

func calculate (revenue float64, expenses float64, taxRate float64) (float64, float64, float64)  {
	var ebt = revenue - expenses
	var profit = ebt * (1- taxRate/100)
	var ratio = ebt /profit
	return ebt, profit, ratio
}

