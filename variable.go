package main

import "fmt"

func main() {
	var name string = "HELLO"
	fmt.Println(name)
	name = "Muhammad Pauzi"
	fmt.Println(name)

	var age = 18

	fmt.Println(age)

	message := "Welcome..." // short way to decla variable

	fmt.Println(message)

	var (
		firstName = "Muhammad"
		lastName  = "Pauzi"
	)

	fmt.Println(firstName, lastName)
}
