package main

import "fmt"

func main() {
	name := "pauzi"

	switch name {
	case "zi":
		fmt.Println("zi")
	case "pauzi":
		fmt.Println("pauzi")
	default:
		fmt.Println("default")
	}

	switch length := len(name); length > 4 {
	case true:
		fmt.Println("true")
	case false:
		fmt.Println("false")
	}

	length := len(name)

	switch {
	case length > 2:
		fmt.Println("> 2")
	case length > 3:
		fmt.Println("> 3")
	case length > 5:
		fmt.Println("> 5")
	default:
		fmt.Println("default")
	}
}
