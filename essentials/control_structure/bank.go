package main

import (
	"fmt"
	"os"
	"strconv"
	"errors"
)

var balance float64;

func main() {
	balance = readBalanceFromFile();
	fmt.Println("Welcome to the Bank of Golang")
	fmt.Println("Please enter your choice")

	fmt.Print("1. Deposit\n2. Withdraw\n3. Balance\n4. Exit\n")

	var choice int

	_, err := fmt.Scan(&choice)
	if err != nil {
		fmt.Println("Error: ", err)
		main()
	}

	choiceOptions(choice)

}

func deposit() {
	var depositAmount float64
	fmt.Print("You have chosen to deposit\nEnter your amount below\n")

	_, err := fmt.Scan(&depositAmount)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	if depositAmount < 0 {
		fmt.Println("Invalid amount")
		choiceOptions(1)
	} else {
		fmt.Println("You have deposited: ", depositAmount)
		balance += depositAmount
	}
	main()
}

func withdraw() {
	var withdrawAmount float64
	fmt.Print("You have chosen to withdraw\nEnter your amount below\n")

	_, err := fmt.Scan(&withdrawAmount)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	if withdrawAmount < 0 {
		err := errors.New("invalid amount")
		fmt.Println("Error: ", err)
		choiceOptions(2)
	} else {
		if withdrawAmount > balance {
			fmt.Println("Insufficient funds")
			choiceOptions(2)
		}
		fmt.Println("You have withdrawn: ", withdrawAmount)
		balance -= withdrawAmount
	}

	main()

}

func checkBalance() {
	fmt.Println("Your balance is: ", balance)
	main()
}

func choiceOptions(choice int) {
	switch choice {
	case 1:
		deposit()
	case 2:
		withdraw()
	case 3:
		checkBalance()
	case 4:
		fmt.Println("You have chosen to exit")
		writeBalanceToFile()
		os.Exit(0)
	default:
		fmt.Println("Invalid choice, please try again")
		main()
	}
}

func writeBalanceToFile() {
	balanceText := fmt.Sprintf("%.2f", balance)
	err := os.WriteFile("balance.txt", []byte(balanceText), 0644)
	if err != nil {
		fmt.Println("Error: ", err)
		panic(err)
	}

}

func readBalanceFromFile() float64 {
	data, err := os.ReadFile("balance.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		return 0;
	}

	currentBalance := string(data)
	if currentBalance == "" {
		return 0;
	}

	balance, err = strconv.ParseFloat(currentBalance, 64)
	if err != nil {
		fmt.Println("Error: ", err)
		return 0;
	}
	return balance
}
