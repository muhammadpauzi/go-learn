package main

import "fmt"

func sayHello() {
	fmt.Println("Hello World w/ function")
}

func main() {
	sayHello()
	WithParam("muh", "pauzi")
	fmt.Println(WithReturn("pauzi"))
	// _, age := WithMultiReturn()
	name, age := WithMultiReturn()
	fmt.Println(name, "umur", age)
	firstname, middlename, lastname := WithNamedReturn()
	fmt.Println(firstname, middlename, lastname)
}

func WithParam(firstname string, lastname string) {
	fmt.Println("Hello", firstname, lastname)
}

func WithReturn(name string) string {
	return "Hello " + name
}

func WithMultiReturn() (string, uint8) {
	return "pauzi", 18
}

func WithNamedReturn() (firstname, middlename, lastname string) {
	firstname = "muh"
	middlename = "-"
	lastname = "pauzi"

	// return firstname, middlename, lastname
	return
}
