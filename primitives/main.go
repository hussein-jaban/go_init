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

		// Basic arithmetic operations built into Golang
				// Only Integers of the same type can be arithmetic operations be done

		fmt.Println(8 + 9)
		fmt.Println(10 - 9)
		fmt.Println(8 * 9)
		fmt.Println(8/2)
		fmt.Println(10 % 3)

		// The Bit Operators
		a := 10
		b := 3
		fmt.Println (a & b) // 0010
		fmt.Println(a | b) // 1011
		fmt.Println(a^ b) // 1001
		fmt.Println(a &^ b)// 0100

		// Bit Shifting
		f := 8 // 2^3
		fmt.Println (f << 3) // 2^3* 2^3 = 2^6
		fmt.Println(f >> 3) // 2^3 / 2^3 = 2^0


}