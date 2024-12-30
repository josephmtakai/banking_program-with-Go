package main

import (
	"fmt"
)

// Function to show current balance
func showBalance(balance float64) {
	fmt.Printf("Your current balance is: $%.2f\n", balance)
}

// Function to deposit money into the account
func deposit(balance float64) float64 {
	var depositAmount float64
	fmt.Print("Enter the amount to deposit: $")
	_, err := fmt.Scan(&depositAmount)
	if err != nil || depositAmount <= 0 {
		fmt.Println("Invalid deposit amount. Please try again.")
		return balance
	}
	balance += depositAmount
	fmt.Printf("$%.2f has been deposited. Your new balance is $%.2f\n", depositAmount, balance)
	return balance
}

// Function to withdraw money from the account
func withdraw(balance float64) float64 {
	var withdrawAmount float64
	fmt.Print("Enter the amount to withdraw: $")
	_, err := fmt.Scan(&withdrawAmount)
	if err != nil || withdrawAmount <= 0 || withdrawAmount > balance {
		fmt.Println("Invalid withdraw amount or insufficient funds. Please try again.")
		return balance
	}
	balance -= withdrawAmount
	fmt.Printf("$%.2f has been withdrawn. Your new balance is $%.2f\n", withdrawAmount, balance)
	return balance
}

func main() {
	var balance float64
	var userChoice int

	for {
		// Display menu options
		fmt.Println("\n*********************")
		fmt.Println("   Banking Program   ")
		fmt.Println("*********************")
		fmt.Println("1. Show Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")
		fmt.Println("*********************")
		fmt.Print("Enter your choice (1-4): ")
		_, err := fmt.Scan(&userChoice)
		if err != nil {
			fmt.Println("Invalid choice, please try again.")
			continue
		}

		switch userChoice {
		case 1:
			showBalance(balance)
		case 2:
			balance = deposit(balance)
		case 3:
			balance = withdraw(balance)
		case 4:
			fmt.Println("Exiting the program. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please choose between 1 and 4.")
		}
	}
}
