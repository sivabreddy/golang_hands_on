package main

// if importing more than one package, use parentheses and write one package per line
// dont use comma or semicolon 
import (
	"fmt"
	"math"
	//"strings"
)
// use camel case for variable names
// go is statically typed programming language
// Go types and Null values
// int - a number without a decimal point
// uint - an unsigned integer (no negative numbers)
// int32 - a 32-bit integer with a range of -2,147,483,648 to 2,147,483,647
// uint32 - an unsigned 32-bit integer with a range of 0 to 4,294,967,295
// int64 - a 64-bit integer with a range of -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807
// float64 - a number with a decimal point
// string - a sequence of characters created via double quotes or backticks
// bool - true or false
// Null values: float64 variable will have a default value of 0.0, bool variable will have a default value of false, string variable will have a default value of an empty string, and int variable will have a default value of 0.
func main() {
	// you can explicitly declare the type of a variable eg: var investmentAmount float64 = 1000
	var investmentAmount float64 = 1000
	var interestRate float64 = 5.5
	var years float64 = 10
	var futureValue = investmentAmount * math.Pow(1+interestRate/100, years)
	fmt.Println("The future value of the investment is:", futureValue)
	// var investmentAmount = 1000
	// var interestRate = 5.5
	// var years = 10
	// var futureValue = float64(investmentAmount) * math.Pow(1+interestRate/100, float64(years))
	// fmt.Println("The future value of the investment is:", futureValue)
}