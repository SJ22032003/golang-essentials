package main

import (
	"fmt"
	"math"
)

const inflationRate float64 = 2.5 // const is a keyword to declare a constant variable
// global variables are declared outside of the main function

func calculateFutureValue(invAmount, years, expectedReturnRate float64) float64 {
	return invAmount * math.Pow(1+expectedReturnRate/100, years)
}

func calculateFutureRealValue(futureValue, years float64) float64 {
	return futureValue / math.Pow(1+inflationRate/100, years)
}

func main() {

	var investmentAmount float64      // var is a keyword to declare a variable,
	years := 10.0
	expectedReturnRate := 5.5 // := is a shorthand for var expectedReturnRate float64 = 5.5, should be used as much as possible

	fmt.Println("Investment Calculator", investmentAmount)

	fmt.Print("Enter investment amount: ")
	_, err := fmt.Scan(&investmentAmount, &years, &expectedReturnRate)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	futureValue := calculateFutureValue(investmentAmount, years, expectedReturnRate)
	futureRealValue := calculateFutureRealValue(futureValue, years)

	fmt.Printf("Future value: %.2f\n", futureValue)  
	fmt.Printf("Future real value: %.2f\n", futureRealValue)

}
