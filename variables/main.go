package main

import (
	"fmt"
	"strconv"
)

// Declaring variables outside scope of main function
// You can't use the shorthand notation
var name string = "john"

// Declaring multiple variables at once
var (
		age int64 = 56
		gender string = "M"
)

// Variables that are Start with a capital letter a exported by go to be used at the package level
var Sentence string = "John is a dude"

func main()  {

	// Declaring a variable without assigning
	var x int

// Assigning the same variable 
	x = 15

	// Declaring a variable and assigning at the same time
	var y int = 34

	// Declaring a variable using the shorthand notation
	z := 45
	fmt.Printf("x: %v, %T\n ", x, x)
	fmt.Printf("y: %v\n ", y)
	fmt.Printf("z: %v\n ", z)
	fmt.Printf("Name: %v\n ", name)

	// Converting int to string
	var i int = 42
	fmt.Printf("%v, %T\n", i, i)
	var j string
	j = strconv.Itoa(i)
	fmt.Printf("%v, %T\n", j, j)
}