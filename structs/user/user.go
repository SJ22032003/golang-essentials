package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func (u User) PrintUserInfo() {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

// This method will clear the user's information we use pointer receiver
func (u *User) ClearUserInfo() {
	u.firstName = ""
	u.lastName = ""
	u.birthdate = ""
	u.createdAt = time.Time{}
}

func New(firstName, lastName, birthdate string) (*User, error) {

	if firstName == "" {
		return nil, errors.New("first name is required")
	}

	if lastName == "" {
		return nil, errors.New("last name is required")

	}

	if birthdate == "" {
		return nil, errors.New("birthdate is required")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}
