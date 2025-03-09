package main

import (
	"fmt"
	"math"
)
// functions are code demand blocks
// main is a function
// 2 types of functions - user defined and built in functions in go libraries
// "func" is keyword used to define a function

// eg: 
// func main() {
// 	outputText()
// }

// func outputText() {
// 	// function body
// 	fmt.Println("Hello, World!")
// }

// Note:
// scope of the variables defined inside a function is limited to that function only
// scope of the variables defined outside a function is global, and is accessible to all the functions 
const inflationRate = 2.5

func main() {
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	// fmt.Print("Investment Amount: ")
	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	// fmt.Print("Expected Return Rate: ")
	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	// fmt.Print("Years: ")
	outputText("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)
	// futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (adjusted for Inflation): %.1f\n", futureRealValue)
	// outputs information
	// fmt.Println("Future Value:", futureValue)
	// fmt.Printf(`Future Value: %.1f\n

	// Future Value (adjusted for Inflation): %.1f`, futureValue, futureRealValue)
	// fmt.Println("Future Value (adjusted for Inflation):", futureRealValue)
	fmt.Print(formattedFV, formattedRFV)
}

func outputText(text string) {
	fmt.Print(text)
}
// return multiple values
func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv := fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}

// alternate approach to return values, its good practice to use return <values>
// func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
// 	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
// 	rfv = fv / math.Pow(1+inflationRate/100, years)
// 	// return fv, rfv
// 	return
// }

// output
// project_7_functions % ./functions
// Investment Amount: 10000
// Expected Return Rate: 5.5
// Years: 10
// Future Value: 17081.4
// Future Value (adjusted for Inflation): 13344.0
