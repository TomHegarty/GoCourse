package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("------------------------")
		panic(err)
	}

	var choice int

	fmt.Println("Welcome to Go Bank")
	fmt.Println("Reach us anytime ", randomdata.PhoneNumber())

	for {
		presentOptions()

		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Current balance: ", accountBalance)
		case 2:
			fmt.Println("Deposit amount: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("invalid amount, must be greater than 0")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Balance updated, current balance: ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)

		case 3:
			fmt.Println("Withdrawal amount: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("invalid amount, must be greater than 0")
				continue
			}

			if accountBalance < withdrawalAmount {
				fmt.Println("insufficient balance")
				continue

			}

			accountBalance -= withdrawalAmount
			fmt.Println(withdrawalAmount, "withdrawn, current balance: ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)

		default:
			fmt.Println("goodbye")
			fmt.Println("thanks for visiting Go bank")
			return
		}
	}
}
