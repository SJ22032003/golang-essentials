package main

func main() {
	age := 32

	if checkIfUserIsAdult(&age) {
		println("User is an adult")
	} else {
		println("User is not an adult")
	}

}

func checkIfUserIsAdult(age *int) bool {
	return *age >= 18
}
