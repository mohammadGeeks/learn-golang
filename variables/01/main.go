package main

import "fmt"

func main() {
	var age int = 20
	fmt.Println(age) //20

	var d = true
	fmt.Println(d) // true

	var age2 = 30
	fmt.Println(age2) //30

	//The := syntax is shorthand for declaring and initializing a variable,
	age3 := 40
	fmt.Println(age3) // 40

	f := "apple"
	fmt.Println(f) // "apple"

	//Variables declared without a corresponding initialization are zero-valued.
	//For example, the zero value for an int is 0.
	var e int
	fmt.Println(e) // 0

	var b, c int = 1, 2
	fmt.Println(b, c) // 1 2
}
