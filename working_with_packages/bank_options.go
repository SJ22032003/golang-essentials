package main

import (
	"errors"
	"fmt"
	"github.com/SJ22032003-bank-packages/balanceFileOps"
	"os"
)

func choiceOptions(choice int, balance *float64) {
	switch choice {
	case 1:
		deposit(balance)
	case 2:
		withdraw(balance)
	case 3:
		checkBalance(balance)
	case 4:
		fmt.Println("You have chosen to exit")
		os.Exit(0)
	default:
		fmt.Println("Invalid choice, please try again")
		main()
	}
}

func deposit(balance *float64) {
	var depositAmount float64
	fmt.Print("You have chosen to deposit\nEnter your amount below\n")

	_, err := fmt.Scan(&depositAmount)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	if depositAmount < 0 {
		fmt.Println("Invalid amount")
		choiceOptions(1, balance)
	} else {
		fmt.Println("You have deposited: ", depositAmount)
		*balance += depositAmount
		balanceFileOps.WriteBalanceToFile(balance)
	}
	main()
}

func withdraw(balance *float64) {
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
		choiceOptions(2, balance)
	} else {
		if withdrawAmount > *balance {
			fmt.Println("Insufficient funds")
			choiceOptions(2, balance)
		}
		fmt.Println("You have withdrawn: ", withdrawAmount)
		*balance -= withdrawAmount
		balanceFileOps.WriteBalanceToFile(balance)
	}

	main()

}

func checkBalance(balance *float64) {
	fmt.Println("Your balance is: ", *balance)
	main()
}
