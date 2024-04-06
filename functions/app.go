package main

import (
	"fmt"
	transform "github.com/SJ22032003/go-functions/transform"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	anonymousTripleFunc := func(num int) int {
		return 3 * num
	}

	tripleNumbers := transform.TransformNumbers(transform.TransformArgs{Numbers: &numbers, Transform: anonymousTripleFunc})
	doubleNumbers := transform.TransformNumbers(transform.TransformArgs{Numbers: &numbers, Transform: transform.Double})

	fmt.Println("Original numbers:", numbers)
	fmt.Println("Triple numbers:", tripleNumbers)
	fmt.Println("Double numbers:", doubleNumbers)


	// Closure
	coinGameAlex := coinGame()
	coinGameAlex()
	coinGameAlex()
	coinGameAlex()

	coinGameBob := coinGame()
	coinGameBob()
	coinGameBob()

	// Variadic function
	fmt.Println("Sum", sum([]int{6, 7, 8, 9, 10}...))

}

// variadic function
func sum(numbers ...int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

// example of closure
func coinGame() func() {
	coin := 5
	return func() {
		coin--
		fmt.Println("Coin:", coin)
	}
}
