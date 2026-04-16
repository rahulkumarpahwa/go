package main

import (
	"fmt"
	"github.com/fatih/color"

	"github.com/rahulkumarpahwa/go/auth"
	"github.com/rahulkumarpahwa/go/user"
)

func main() {
	auth.LoginWithCredentials("apple", "password")

	user := user.User{
		Email:    "apple@apple.com",
		Password: "password",
		Username: "Apple",
	}
	fmt.Println(user)

	color.Green("this is the string")

}
