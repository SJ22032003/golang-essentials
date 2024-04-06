package main

import "fmt"

func main() {
	prices := []int{1, 2, 3, 4, 5} // slice
	updatedPrices := append(prices, 6)

	updatedPrices =  append(updatedPrices, prices...)

	fmt.Println(prices, updatedPrices)

}

// func main() {
// 	prices := [5]int{1, 2, 3, 4, 5}

// 	featuredPrice := prices[:3]

// 	featuredPrice[0] = 100

// 	fmt.Println(len(prices), cap(prices))

// }
