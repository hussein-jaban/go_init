package main

import "fmt"

// Declaring variables outside scope of main function
// You can't use the shorthand notation
var name string = "john"

// Declaring multiple variables at once
var (
		age int64 = 56
		gender string = "M"
)

// 

func main()  {

	// Declaring a variable without assigning
	var x int

// Assigning the same variable 
	x = 15

	// Declaring a variable and assigning at the same time
	var y int = 34

	// Declaring a variable using the shorthand notation
	z := 45
	fmt.Printf("x: %v, %T ", x, x)
	fmt.Printf("y: %v ", y)
	fmt.Printf("z: %v ", z)
	fmt.Printf("Name: %v ", name)
}