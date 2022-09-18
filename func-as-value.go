package main

import "fmt"

func sayGoodBye() string {
	return "goodbyye"
}

func main() {
	goodBye := sayGoodBye
	fmt.Println(goodBye())
}
