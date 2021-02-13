package main

import (
	"fmt"
	"strings"
)

func main() {
	m1 := "hello world"
	m2 := "wo"

	fmt.Println(strings.Contains(m1, m2))
	fmt.Println(strings.ReplaceAll(m2, "w", "G"))
}
