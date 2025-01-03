package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/nikhil/podcast/auth"
	"github.com/nikhil/podcast/user"
)

func main() {
	auth.LoginWithCredentials("codersgyan", "pass")

	session := auth.GetSession()
	fmt.Println(session)

	user := user.User{
		Email: "user@gmail.com",
		Name:  "john doe",
	}

	// fmt.Println(user.Email, user.Name)
	color.Red(user.Email)

}
