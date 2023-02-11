package main

import (
	"fmt"
)

// Primitive Data Types

func main()  {

	// Booleans Types
	var	j bool = true
	fmt.Printf("%v, %T\n", j, j) // true, bool

	// When initialization a variable without assign a value, go assigns a zero value which is falsly in nature.
	var c bool
		fmt.Printf("%v, %T\n", c, c) // false, bool

		// Numeric Types in Go

		// Integers
		var i int = 1
  fmt.Printf("%v, %T\n", i, i) // 1, int
		var q int8 = 1
  fmt.Printf("%v, %T\n", q, q) // 1, int8
		var w int32 = 1
  fmt.Printf("%v, %T\n", w, w) // 1, int32
		var z int64 = 1
  fmt.Printf("%v, %T\n", z, z) // 1, int64

		// Basic arithmetic operations built into Golang
				// Only Integers of the same type can be arithmetic operations be done

		fmt.Println(8 + 9) // 17
		fmt.Println(10 - 9) // 1
		fmt.Println(8 * 9) // 72
		fmt.Println(8/2) // 4
		fmt.Println(10 % 3) // 1

		// The Bit Operators
		a := 10
		b := 3
		fmt.Println (a & b) // 0010 =  2
		fmt.Println(a | b) // 1011  = 11
		fmt.Println(a^ b) // 1001 =  9
		fmt.Println(a &^ b)// 0100 =  8

		// Bit Shifting
		f := 8 // 2^3
		fmt.Println (f << 3) // 2^3* 2^3 = 2^6 = 64
		fmt.Println(f >> 3) // 2^3 / 2^3 = 2^0 = 1


		// Floats ie float32, float64
		var n float64 = 3.14
		n = 13.7e72
		n = 2.1E14
		fmt. Printf("%v, %T\n", n, n) // 2.1e+14, float64

		ab := 10.2
		bb := 3.7
		fmt.Println(ab + bb)  // 13.899999999999999
		fmt.Println(ab - bb)  // 6.499999999999999
		fmt.Println(ab * bb)  // 37.74
		fmt.Println(ab / bb)  // 2.7567567567567566


	// Strings
		s := "this is a string"
		fmt. Printf("%v, %T\n", s, s) // this is a string, string

		// string concatination
		s2 := "This also is string two"
		fmt.Println(s + s2) // this is a stringThis also is string two
		by := []byte(s)
		fmt. Printf("%v, %T\n", by, by) // [116 104 105 115 32 105 115 32 97 32 115 116 114 105 110 103], []uint8

}