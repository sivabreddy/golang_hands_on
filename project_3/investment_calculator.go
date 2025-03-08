package main

import (
	"fmt"
	"math"
)

func main() {
	// constants are immutable - we can't change the value of constant
	const inflationRate float64 = 2.5
	// variables can be reassigned, i.e we can change the value of variable - mutable
	// when we don't define explicit type, we can use := operator
	// its a good practise to use this approach
	// we can add the variable of same type in one line with comma separated values
	var investmentAmount , years float64 = 1000 , 10
	interestRate := 5.5
	futureValue := investmentAmount * math.Pow(1+interestRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Println("The future value of the investment is:", futureValue)
	fmt.Println("The future real value of the investment is:", futureRealValue)
}