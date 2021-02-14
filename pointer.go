package main

import "fmt"

func main() {
	x := 10
	a := 10
	b := 20
	x = doubleNor(x)
	fmt.Println(x)
	doublePtr(&x)
	fmt.Println(x)
	func(x *int, y *int) {
		*x += *y
		*y = *x - *y
		*x = *x - *y
	}(&a, &b)
	fmt.Printf("a: %d, b: %d", a, b)
}

func doublePtr(x *int) {
	*x = *x * 2
}

func doubleNor(x int) int {
	return x * 2
}
