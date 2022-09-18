package main

import "fmt"

func sumAll(numbers ...int) int {
	fmt.Println(numbers)
	total := 0
	for _, v := range numbers {
		total += v
	}
	return total
}

func main() {
	total := sumAll(10, 20, 30, 40, 50)
	fmt.Println(total)

	slice := []int{10, 21, 43, 546, 12}
	total2 := sumAll(slice...) // spread
	fmt.Println(total2, slice)
}
