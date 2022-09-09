package main

import "fmt"

func main() {
	const firstName string = "Muhammad"
	const lastName = "Pauzi"

	const age uint8 = 18

	const (
		var1 string = "VAR1"
		var2        = "VAR2"
	)

	fmt.Println(firstName, lastName)
	fmt.Println(var1, var2)
	fmt.Println(age)
}
