package main

import (
	"fmt"
)

const a int16 = 27

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
} 