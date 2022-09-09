package main

import "fmt"

func main() {
	type NoKTP string

	var ktp NoKTP = "12345678"

	fmt.Println(ktp)
}
