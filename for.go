package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println(counter, "counter")
		counter++
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i, "index")
	}

	names := []string{"zi", "oji", "pauzi", "muhammad"}

	for i, name := range names {
		fmt.Println(name, i)
	}
}
