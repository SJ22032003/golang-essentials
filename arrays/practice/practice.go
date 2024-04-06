package main

import (
	"fmt"

	productStruct "github.com/SJ22032003/array-practice-1/struct"
)

func main() {
	product1 := productStruct.New("Laptop", 1, 1000.0)
	product2 := productStruct.New("Mouse", 2, 20.0)

	products := []productStruct.Product{*product1, *product2}

	products = append(products, *productStruct.New("Keyboard", 3, 50.0))

	fmt.Println(products)

	i := 0
	for i < len(products) {
		fmt.Println(products[i].Title, products[i].Id, products[i].Price)
		i++
	}


}

// func main() {

// 	hobbies := [3]string{"Reading", "Coding", "Gaming"}
// 	fmt.Println(hobbies)
// 	fmt.Println(hobbies[0])
// 	fmt.Println((hobbies[1:]))

// 	// tryArr := [2]string{hobbies[0], hobbies[1]} // first way
// 	// fmt.Println(tryArr)

// 	updatedHobbies := hobbies[:2]
// 	fmt.Println(updatedHobbies)

// 	updatedHobbies = hobbies[1:]
// 	fmt.Println(updatedHobbies)

// 	courseGoals := []string{"Learn Go", "Learn Python"}
// 	courseGoals[1] = "Learn Java"
// 	courseGoals = append(courseGoals, "Learn C++")
// 	fmt.Println(courseGoals)

// }

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
