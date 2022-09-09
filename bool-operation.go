package main

import "fmt"

func main() {
	fmt.Println(true && true, true || false, false || false, false && true, (!true) && false, false && false)
}
