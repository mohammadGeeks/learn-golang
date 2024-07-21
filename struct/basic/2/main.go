package main

import "fmt"

type Person struct {
	First string
	Last  string
}

func main() {
	p1 := NewPerson("James", "Bond")
	cop := p1
	cop.Last = "test"

	fmt.Println(cop)
	fmt.Println(p1)

	p2 := NewPerson2("lebron", "james")
	instance := p2

	instance.First = "james"
	instance.Last = "harden"

	fmt.Println(instance)
	fmt.Println(p2)

}

func NewPerson(f, l string) *Person {
	np := Person{f, l}
	return &np
}

func NewPerson2(f, l string) Person {
	np := Person{f, l}
	return np
}
