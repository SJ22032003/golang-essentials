package main

import (
	"fmt"
	// "github.com/Pallinder/go-randomdata"
	"github.com/SJ22032003-bank-packages/balanceFileOps"
)

func main() {
	var balance float64 = balanceFileOps.ReadBalanceFromFile()
	fmt.Println("Welcome to the Bank of Golang")
	fmt.Println("Please enter your choice")

	fmt.Print("1. Deposit\n2. Withdraw\n3. Balance\n4. Exit\n")

	var choice int

	_, err := fmt.Scan(&choice)
	if err != nil {
		fmt.Println("Error: ", err)
		main()
	}

	// fmt.Println("You can reach us at", randomdata.PhoneNumber())
	choiceOptions(choice, &balance)


}
