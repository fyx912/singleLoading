package main

import (
	"fmt"
	"singleLoading/models"
)

func main() {
	models.ValidationLogin()
	var user models.User
	user.Username = "ding"
	user.Password = "123456"
	fmt.Println(user)
}
