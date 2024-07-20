package main

import "fmt"

func main() {
	// In Go, variables that are declared but not initialized are automatically assigned a "zero value.

	var a int    // a is 0
	var b bool   // b is false
	var c string // c is ""
	var d *int   // d is nil

	fmt.Println(a, b, c, d) // 0 false  <nil>
}
