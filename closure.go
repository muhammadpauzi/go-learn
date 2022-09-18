package main

import "fmt"

func main() {
	counter := 0
	name := "pauzi"
	increment := func() {
		fmt.Println("incrementing...")
		name := "in clouse scope"
		fmt.Println(name)
		counter++
	}

	increment()
	increment()
	fmt.Println(name)
	fmt.Println(counter, "counter")
}
