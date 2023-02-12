package main

import (
	"fmt"
)

const a int16 = 27

// Using iota with constants
const (
				x = iota
				y
				z
)

// Iota resets to 0
const (
			x2 = iota
)

func main()  {
	const myConst int = 23
	fmt.Printf("%v, %T\n", myConst, myConst) // 23, int

	// Shadowing
	const a int = 14
 fmt.Printf("%v, %T\n", a, a) // 14, int

	// Constants type inference
	const c = 42
	var b int16 = 27
	fmt. Printf("%v, %T\n", c + b, c + b) // 69, int16

	fmt.Printf("%v\n", x) // 0
	fmt.Printf("%v\n", y) // 1
	fmt.Printf("%v\n", z) // 2
	fmt.Printf("%v\n", x2) // 0
} 