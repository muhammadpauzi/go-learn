package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}

		fmt.Println(i)
	}

	fmt.Println("done!")

	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println("genap:", i)
	}

	fmt.Println("done!")

}
