package adminStruct

import (
	"errors"
	"fmt"
	"github.com/SJ22032003-struct/user"
)

type Admin struct {
	email    string
	password string
	user.User
}

type AdminArgs struct {
	Email     string
	Password  string
	FirstName string
	LastName  string
	Birthdate string
}

func (a Admin) PrintAdminInfo() {
	fmt.Println(a.email, a.password)
	a.PrintUserInfo()
}

func New(args AdminArgs) (*Admin, error) {

	if args.Email == "" {
		return nil, errors.New("email is required")
	}

	if args.Password == "" {
		return nil, errors.New("password is required")
	}

	user, error := user.New(args.FirstName, args.LastName, args.Birthdate)
	if error != nil {
		return nil, error
	}

	return &Admin{
		email:    args.Email,
		password: args.Password,
		User:     *user,
	}, nil
}
