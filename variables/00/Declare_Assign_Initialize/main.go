package main

import "fmt"

func main() {

	//Declaration is the process of telling the Go compiler that you intend to use a variable by specifying its name and type.
	//This does not assign any value to the variable; it merely allocates space for it. For example:
	var x int // Declaration without initialization

	//Assignment is the act of setting a value to a variable that has already been declared.
	//The assignment operator = is used for this purpose.
	//You can assign a value to a variable either at the time of declaration or later in the program.

	x = 42 // Assignment

	fmt.Println(x) // 42

	//Initialization is the act of declaring a variable and assigning it a value simultaneously.
	//This can be done in a variety of ways in Go:
	var m int = 42 // Declaration with initialization
	var y = 42     // Type inference
	z := 42        // Short variable declaration and initialization

	fmt.Println(z, y, m) //42 42 42
}
