package main

import "fmt"

func main() {
	if 1 == 1 {
		fmt.Println("MASUK")
	}

	if 1 == 2 {
		fmt.Println("MASUK")
	} else {
		fmt.Println("MASUK ELSE")
	}

	if 1 == 2 {
		fmt.Println("MASUK")
	} else if 1 == 1 {
		fmt.Println("MASUK 2")
	} else {
		fmt.Println("MASUK ELSE")
	}

	if a := len("asd"); a > 2 {
		fmt.Println("MASUK")
	}
}
