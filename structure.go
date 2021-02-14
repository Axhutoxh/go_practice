package main

import "fmt"

type Car struct {
	Name     string
	Age      int
	MobileNo int
}

func (c Car) Print() {
	fmt.Println(c)
}

func main() {
	c := Car{"Axhutoxh", 23, 9455118809}
	c.Print()
	fmt.Printf("Name = %s Age = %d MobileNo = %d ", c.Name, c.Age, c.MobileNo)
}
