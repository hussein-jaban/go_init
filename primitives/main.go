package main

import (
	"fmt"
)

// Primitive Data Types

func main()  {

	// Booleans Types
	var	j bool = true
	fmt.Printf("%v, %T\n", j, j)

	// When initialization a variable without assign a value, go assigns a zero value which is falsly in nature.
	var c bool
		fmt.Printf("%v, %T\n", c, c)

		// Numeric Types in Go

		// Integers
		var i int = 1
  fmt.Printf("%v, %T\n", i, i)
		var q int8 = 1
  fmt.Printf("%v, %T\n", q, q)
		var w int32 = 1
  fmt.Printf("%v, %T\n", w, w)
		var z int64 = 1
  fmt.Printf("%v, %T\n", z, z)

		

}