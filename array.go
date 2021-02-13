package main

import "fmt"

func todo() {
	//var arr []int
	arr := []int{1, 2, 3, 4}
	arr2 := []string{"hi", "hello", "word"}
	fmt.Println(arr)
	fmt.Println(arr2)
	arr2 = append(arr2, "its world")
	fmt.Println(arr2)
}

func main() {
	todo()
}
