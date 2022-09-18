package main

import "fmt"

// defer : function yang selalu berjalan mesikipun program/function yang mengandung defer terjadi error

func logging() {
	fmt.Println("logging...")
}

func runApplication(value int) {
	defer logging()
	fmt.Println("run app...", 10/value)
}

func main() {
	runApplication(1)
	// runApplication(0) // error
}
