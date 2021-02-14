package main

import "fmt"

type Car interface {
	Drive()
	Stop()
}

type Eco struct {
	EcoType string
}

type Luxury struct {
	LuxuryType string
}

func NewModel(arg string) Car {
	return &Luxury{arg}
}

func (l *Luxury) Drive() {
	fmt.Println("Luxury Starts")
	fmt.Println(l.LuxuryType)
}

func (e *Eco) Drive() {
	fmt.Println("Eco Starts")
	fmt.Println(e.EcoType)
}

func (l *Luxury) Stop() {
	fmt.Println("BMW Stops")
}

func main() {
	l := Luxury{"BMW"}
	e := Eco{"Maruti"}
	l.Drive()
	e.Drive()
	l.Stop()
}
