package main

import "fmt"

func main() {

	//three categories of types in Go
	//builtin: numeric, string, bool, array
	//reference: slices, maps, channels, pointers
	//struct

	//1 - Built-in Types:
	// Built-in types are types that are part of the Go language itself and provide the basic building blocks for constructing more complex types.
	// Some commonly used built-in types include:

	// Numeric types (int, float64, etc.)
	// Boolean type (bool)
	// String type (string)

	var i int = 10
	var f float64 = 3.14
	var b bool = true
	var s string = "hello"

	println(i, f, b, s)

	// 2- Reference Types:
	//Reference types are types that hold references to the underlying data.
	//They can be thought of as "containers" that point to the actual value or a collection of values.
	//Reference types in Go include:
	//Slices
	//Maps
	//Channels
	//Pointers

	// Slice
	nums := []int{1, 2, 3}

	// Map
	mapping := map[string]int{
		"one": 1,
		"two": 2,
	}

	// Channel
	ch := make(chan int)

	// Pointer
	var ptr *int

	fmt.Println(nums, mapping, ch, ptr)

	// 3- Struct Types:
	//Struct types are composite data types that group variables under a single name.
	//These variables can be of different types, and each field in the struct is given a name.
	//Structs are used for grouping data and are particularly useful when you want to create complex types that aggregate multiple pieces of information.

	type Person struct {
		Name    string
		Age     int
		Address string
	}

	var p Person
	p.Name = "John"
	p.Age = 30
	p.Address = "123 Main St"
	fmt.Println(p)
}
