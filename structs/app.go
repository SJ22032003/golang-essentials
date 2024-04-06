package main

import (
	"fmt"
	"github.com/SJ22032003-struct/admin"
	"github.com/SJ22032003-struct/user"
	"time"
)

func main() {
	println(time.April)
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser, error := user.New(firstName, lastName, birthdate)
	if error != nil {
		fmt.Println(error)
		return
	}

	appUser2, error := user.New(firstName, lastName, birthdate)
	if error != nil {
		fmt.Println(error)
		return
	}

	adminArgs := adminStruct.AdminArgs{
		Email:     "alex@gmail.com",
		Password:  "password",
		FirstName: firstName,
		LastName:  lastName,
		Birthdate: birthdate,
	}

	appAdmin, error := adminStruct.New(adminArgs)
	if error != nil {
		fmt.Println(error)
		return
	}

	appAdmin.PrintAdminInfo()

	appUser.PrintUserInfo()
	appUser2.PrintUserInfo()
	appUser.ClearUserInfo()

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
