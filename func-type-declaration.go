package main

import (
	"fmt"
	"strings"
)

type Filter func(string) string

func sayHelloWithFilter(text string, filter Filter) {
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
