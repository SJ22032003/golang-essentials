package balanceFileOps

import (
	"fmt"
	"os"
	"strconv"
)

func WriteBalanceToFile(balance *float64) {
	balanceText := fmt.Sprintf("%.2f", *balance)
	err := os.WriteFile("balance.txt", []byte(balanceText), 0644)
	if err != nil {
		fmt.Println("Error: ", err)
		panic(err)
	}

}

func ReadBalanceFromFile() float64 {
	data, err := os.ReadFile("balance.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		return 0;
	}

	currentBalance := string(data)
	if currentBalance == "" {
		return 0;
	}

	var userBalance float64
	userBalance, err = strconv.ParseFloat(currentBalance, 64)
	if err != nil {
		fmt.Println("Error: ", err)
		return 0;
	}
	return userBalance
}