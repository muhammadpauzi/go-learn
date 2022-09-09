package main

import "fmt"

func main() {
	// pointer, length, capacity
	var months = [...]string{
		"jan", "feb", "mar", "apr", "mei", "jun", "Jul", "agu", "sep", "okt", "nov", "des",
	}

	var slice1 = months[1:4]
	fmt.Println(months)
	slice1[1] = "BERUBAH"
	fmt.Println(slice1)
	fmt.Println(months)
	fmt.Println(len(slice1), cap(slice1))

	slice1 = append(slice1, "Baru") // buat array baru jika array lama tidak cukup capacitynya
	fmt.Println(slice1)
	fmt.Println(months)

	// slice dari awal
	newSlice := make([]string, 10)
	newSlice[0] = "0"
	newSlice[4] = "4"
	fmt.Println(newSlice)

	var a = [...]int{1, 2, 3, 4} // array [...]
	var b = []int{1, 2, 3, 4, 5} // slice []
	fmt.Println(a, b)
}
