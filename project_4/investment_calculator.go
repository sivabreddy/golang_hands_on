package main

import (
	"fmt"
	"math"
)

func main() {
	// constants are immutable - we can't change the value of constant
	// constants cant be passes during runtime
	const inflationRate = 2.5
	// variables can be reassigned, i.e we can change the value of variable - mutable
	// when we don't define explicit type, we can use := operator
	// its a good practise to use this approach
	// we can add the variable of same type in one line with comma separated values
	var investmentAmount float64 // investment amount variable is assigned its default value which is a "0.0" , float64 value.
	var years float64
	var interestRate float64
    // pointer to a variable is a reference to the memory location of the variable
	fmt.Print("Enter the investment amount:\n")
	fmt.Scan(&investmentAmount)
	fmt.Print("Enter the years:\n")
	fmt.Scan(&years)
	fmt.Print("Enter the interest rate:\n")
	fmt.Scan(&interestRate)
    // even this works
    // fmt.Scan(&investmentAmount, &years, &interestRate)
	futureValue := investmentAmount * math.Pow(1+interestRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println("The future value of the investment is:", futureValue)
	fmt.Println("The future real value of the investment is:", futureRealValue)
}