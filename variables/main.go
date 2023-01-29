package main

import "fmt"

func main()  {

	// Declaring a variable without assigning
	var x int
// Assigning the same variable 
	x = 15

	// Declaring a variable and assigning at the same time
	var y int = 34

	// Declaring a variable using the shorthand expression
	z := 45
	fmt.Printf("x: %v ", x)
	fmt.Printf("x: %v ", y)
	fmt.Printf("x: %v ", z)
}