package main

import "fmt"

func main() {
    var accountBalance float64= 1000
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
			fmt.Println("4. Exit")
	
			var choice int
			fmt.Println("Your choice:")
			fmt.Scanln(&choice)
		
		if choice == 1 {
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
			fmt.Println("Your current balance is:", accountBalance)
		} else if choice == 2 {
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
				fmt.Println("Your current balance is:", accountBalance)
			}
		} else if choice == 3 {
			fmt.Println("Your current balance is:", accountBalance)
		} else if choice == 4 {
			fmt.Println("Exit")
		} else {
			fmt.Println("Invalid choice")
			// exit from main
			// return
			break
		}
		}
		fmt.Println("Thanks for choosing our bank")
		fmt.Println("Have a nice day!")
}