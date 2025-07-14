package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const fileName = "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 1000, errors.New("failed to find balance file")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("failed to parse stored balance value")
	}
	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(fileName, []byte(balanceText), 0644)
}

func main() {
	var accountBalance, err = getBalanceFromFile()
	fmt.Println("Welcome to Go Bank!")

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-------------------------")
		// panic(err)
	}

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your account balance is: $", accountBalance)
		case 2:
			fmt.Print("Enter amount to deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Deposit amount must be greater than zero.")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Balance updated! Your new balance is: $", accountBalance)
			writeBalanceToFile(accountBalance)
		case 3:
			fmt.Print("Enter amount to withdraw: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Deposit amount must be greater than zero.")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Insufficient funds! Your current balance is: $", accountBalance)
				continue
			}

			accountBalance -= withdrawAmount
			fmt.Println("Balance updated! Your new balance is: $", accountBalance)
			writeBalanceToFile(accountBalance)
		default:
			fmt.Println("Exiting the bank. Thank you for using Go Bank!")
			fmt.Println("Thanks for choosing Go Bank!")
			return
		}
	}

}
