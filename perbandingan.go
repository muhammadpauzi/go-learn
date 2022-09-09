package main

import "fmt"

func main() {
	fmt.Println(10 < 11, 2 > 1, 1 == 1, 1 != 2, (10+2) > 13)

	fmt.Println("A" > "B", "J" < "K") // using byte
}
