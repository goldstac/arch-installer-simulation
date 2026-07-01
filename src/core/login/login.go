package main

import "fmt"

func main() {
	var username string
	var passwd string
	fmt.Println("Welcome To Arch Linux")
	fmt.Println("Enter Username")
	fmt.Scan(&username)
	if username == "admin" {
		fmt.Println("Username Correct")
		fmt.Println("Enter Password")
		fmt.Scan(&passwd)
		if passwd == "iusearchbtw" {
			fmt.Println("Password successful")
			fmt.Println("Welcome", username)
		} else {
			fmt.Println("Password Incorrect")
		}
	} else {
		fmt.Println("Username Incorrect")
	}

}
