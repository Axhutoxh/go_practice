package main

import "fmt"

func Anything(anything interface{}) {
	fmt.Println(anything)
}

func main() {
	fmt.Println("go")
	Anything(32362)
	Anything("Ashutosh")
	Anything(struct{}{})

	mymap := make(map[string]interface{})

	mymap["name"] = "NAME"
	mymap["age"] = 10

	fmt.Println(mymap)
}
