package main

import "fmt"

func main() {
	fmt.Println(10+10, 10-10, 10*10, 10%2, 10/5)

	// augmented assignment
	var a int32 = 10
	a += 10
	a -= 5
	a *= 2
	a %= 2
	a /= 10
	fmt.Println(a)

	// unary operator
	// ++ == a = a + 1
	var u = -10
	u++
	u++
	u--
	// ! bool
	fmt.Println(u)

}
