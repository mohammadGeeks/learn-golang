package main

import "fmt"

const (
	small = 0
	big   = 2
)

const Pi = 3.14

func main() {
	const x = 42
	fmt.Println(x)

	fmt.Println(small, big)
	fmt.Println(Pi)
}
