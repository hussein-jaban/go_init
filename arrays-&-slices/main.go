package main

import (
	"fmt"
)

func main()  {
	// Arrays
	grades := [3]int{97, 85, 93}
 fmt.Printf("Grades: %v\n", grades) // Grades: [97 85 93]

	// Using the lateral syntax to declare arrays
	grades2 :=[...]int{97, 85, 93}
 fmt.Printf("Grades2: %v\n", grades2) // Grades2: [97 85 93]

	var students [3]string
	fmt.Printf("Students: %v\n", students) // Students: [  ]
	students[0] = "hussein"
	students[1] = "john"
	students[2] = "ali"
	fmt.Printf("Students: %v\n", students) // Students: [hussein john ali]

	// Slices
	z := []int{97, 85, 93}
	fmt.Printf("Z: %v\n", z) // Z: [97 85 93]

	// create slice using the brackests syntax []
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[:] // slice of all elements
	c := a[3:] // slice from 4th element to end
	d := a[:6] // slice first 6 elements
	e := a[3:6] // slice the 4th, 5th, and 6th elements
	fmt.Println(a) // [1 2 3 4 5 6 7 8 9 10]
	fmt.Println(b) // [1 2 3 4 5 6 7 8 9 10]
	fmt.Println(c) // [4 5 6 7 8 9 10]
	fmt.Println(d) // [1 2 3 4 5 6]
	fmt.Println(e) // [4 5 6]
} 