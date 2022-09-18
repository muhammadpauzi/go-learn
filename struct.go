package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           uint8
}

func sayHello(customer Customer, name string) {
	fmt.Println("Hello,", customer.Name, name)
}

func (customer Customer) sayHello2(name string) {
	fmt.Println("Hello,", customer.Name, name)
}

func main() {
	var cus1 Customer
	cus1.Name = "pauzi"
	cus1.Age = 18
	cus1.Address = "ind"
	fmt.Println(cus1)

	sayHello(cus1, "oji")
	cus1.sayHello2("oji")

	cus2 := Customer{
		Name:    "pauzi",
		Address: "ind",
		Age:     18,
	}
	fmt.Println(cus2, cus2.Name)

	cus3 := Customer{"pauzi", "ind", 18}
	fmt.Println(cus3, cus3.Name)
}
