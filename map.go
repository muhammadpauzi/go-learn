package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Pauzi",
		"address": "asasas",
	}

	person["title"] = "Judul"

	fmt.Println(person)
	delete(person, "address")
	fmt.Println(person["name"], person)
}
