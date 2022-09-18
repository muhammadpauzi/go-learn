package main

import "fmt"

func endApp() {
	fmt.Println("end app!")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("my app error") // app akan berhenti, namun defer tetap berjalan
	}
}

func main() {
	runApp(false)
	runApp(true)
}
