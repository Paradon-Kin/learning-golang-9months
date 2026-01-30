package main

import (
	"errors"
	"fmt"
)

var ErrorUsername = errors.New("user not found")
var ErrorPassword = errors.New("wrong password")

type Account struct {
	Username string
	password string
}

func (a *Account) Login(username, password string) (bool, error) {
	if username != a.Username {
		return false, ErrorUsername
	} else if password != a.password {
		return false, ErrorPassword
	} else {
		return true, nil
	}
}

func main() {
	myacc := &Account{Username: "Admin", password: "1234"}
	acc1 := &Account{Username: "sunny", password: "123"}

	_, err := myacc.Login("admin1", "1234")
	_, err1 := acc1.Login("sunny", "123")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Login Success (%s)\n", myacc.Username)
	}

	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Printf("Login Success (%s)\n", acc1.Username)
	}
}
