package main

import "fmt"

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
		// fmt.Println(i)
	}
	return result
}

func factorialRecursive(value int64) int64 {
	if value == 1 {
		return 1
	} else {
		// fmt.Println(value)
		return value * factorialRecursive(value-1)
	}
}

func main() {
	fmt.Println(factorialLoop(10))
	fmt.Println(factorialRecursive(10))
}
