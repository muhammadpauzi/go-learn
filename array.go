package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "'1'"
	names[1] = "'2'"
	names[2] = "'3'"

	fmt.Println(names, names[1])

	var values = [3]int{
		1,
		2,
		3,
	}

	values[2] = 10

	fmt.Println(values, values[1], len(values))

	var a [12]int

	fmt.Println(len(a))
}
