package main

func main() {

	factorialNum := 25
	factorialResult := factorial(factorialNum)
	println("Factorial of", factorialNum, "is", factorialResult)

}

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}
