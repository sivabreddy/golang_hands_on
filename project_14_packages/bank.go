package main

import (
	// packages part of standard library
	"fmt"
	// packages part of other packages
	// Important: always the functions in the other packages when importing need to use the package name in the go mod
	// functions in the other packages must be defined with captial letter
	// when using those functions from other packages we need to use the <package name>.<function name>
	"example.com/packages/fileOps"

	// Third-party packages
	// go package discovery page eg: go-randomdata, you find repository, documentation etc.
	// installation command: go get <package path from github>
	"github.com/Pallinder/go-randomdata"
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



func main() {
		var accountBalance, err = fileOps.GetFloatFromFile(accountBalanceFile)
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
	    fmt.Println("reach us 24/7",randomdata.PhoneNumber())
		
        for {
			
            displayOptions()
	
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
			fileOps.WriteFloatToFile(accountBalance, accountBalanceFile)
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
				fileOps.WriteFloatToFile(accountBalance, accountBalanceFile)
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
