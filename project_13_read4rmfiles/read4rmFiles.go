package main

import (
	"fmt"
	"os"
	"strconv"
	"errors"
)

const accountBalanceFile = "balance.txt"
// this block works if there is already file with balance value in it.
// suppose if there is no balance.txt file the it should return error. so need to handle it.
// read from file
// func readBalanceFromFile() float64 {
// 	fmt.Println("Reading balance from file")
// 	data, _ := os.ReadFile(accountBalanceFile)
// 	balanceText := string(data)
// 	balance, _ := strconv.ParseFloat(balanceText, 64)
// 	return balance
// 	}
// write to file
// func writeBalanceToFile(accountBalance float64) {
// 	fmt.Println("Writing balance to file")
// 	balanceWrite := fmt.Sprintf("%.2f", accountBalance)
// 	// 0644 permisssions
// 	os.WriteFile(accountBalanceFile, []byte(balanceWrite), 0644)
// }

// readfile with error handling
// read from file
func readBalanceFromFile() (float64, error) {
	fmt.Println("Reading balance from file")
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return 0.0, errors.New("error reading file")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		fmt.Println("Error parsing balance:", err)
		return 0.0, errors.New("failed to parse stored balance value")
	}
	return balance, nil
	}

// write to file
func writeBalanceToFile(accountBalance float64) {
	fmt.Println("Writing balance to file")
	balanceWrite := fmt.Sprintf("%.2f", accountBalance)
	// 0644 permisssions
	os.WriteFile(accountBalanceFile, []byte(balanceWrite), 0644)
}


func main() {
		var accountBalance, err = readBalanceFromFile()
		if err != nil {
			fmt.Println("Error reading balance:")
			// if we don't want to continue, we can return here
			// return
			// we can stop using panic()
			// panic("can't continue further")
		}
	    // Important: we have only for loop in go.
		// example to run the code for 2 iterations with an condition.
		// for i:=0; i < 2; i++ {
		// to run in infinite loop we don't need to specify the condition.
		fmt.Println("Welcome to Go Bank")
        for {
			
			fmt.Println("what you want to do?")
			fmt.Println("1. Deposit")
			fmt.Println("2. Withdraw")
			fmt.Println("3. Check Balance")
	
			var choice int
			fmt.Println("Your choice:")
			fmt.Scanln(&choice)

		switch choice {

			case 1:
				fmt.Println("Your Deposit:")
			var deposit float64
			fmt.Scanln(&deposit)
			if deposit < 0 {
				fmt.Println("Invalid deposit, must be greater than 0")
				// naked return, it stops execution.
				// return
				continue
			}
			accountBalance += deposit
			writeBalanceToFile(accountBalance)
			fmt.Println("Your current balance is:", accountBalance)


			case 2:
				fmt.Println("Withdraw")
			var withdraw float64
			fmt.Scanln(&withdraw)
			if withdraw < 0 {
				fmt.Println("Invalid withdraw, must be greater than 0")
				// return
				continue
			}
			if withdraw > accountBalance {
				fmt.Println("Insufficient balance")
			} else {
				accountBalance -= withdraw
				writeBalanceToFile(accountBalance)
				fmt.Println("Your current balance is:", accountBalance)
			}

			case 3:
				fmt.Println("Your current balance is:", accountBalance)

			default:
				fmt.Println("Invalid choice")
				fmt.Println("Thanks for choosing our bank")
		        fmt.Println("Have a nice day!")
				// can't use break here it can't take outside the loop, instead use return
				// break
				return
			}
		}
}