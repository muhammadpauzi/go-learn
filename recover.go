package main

import "fmt"

func endApp() {
	fmt.Println("end app!")

	message := recover() // ambil error dari panic, dan pastikan menjalankan recover func di bagian lain dari panic func scope
	if message != nil {
		fmt.Println("panic error founded:", message)
	}
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
