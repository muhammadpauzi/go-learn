package main

import (
	"fmt"
	"strings"
)

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked (blacklist)", name)
	} else {
		fmt.Println("Welcome")
	}
}

// func blacklist(name string)bool{

// }

func main() {
	registerUser("pauzi", func(s string) bool {
		return strings.Contains(s, "example")
	})

	registerUser("asd example", func(s string) bool {
		return strings.Contains(s, "example")
	})
}
