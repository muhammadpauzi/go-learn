package main

import (
	"fmt"
	"strings"
)

func sayHelloWithFilter(text string, filter func(string) string) {
	filteredText := filter(text)
	fmt.Println("Hello", filteredText)
}

func spamFilter(text string) string {
	if text == "anjing" {
		return "******"
	}

	return text
}

func upperFilter(text string) string {
	return strings.ToUpper(text)
}

func main() {
	sayHelloWithFilter("anjing", spamFilter)
	sayHelloWithFilter("asd", spamFilter)
	sayHelloWithFilter("pauzi", upperFilter)
}
