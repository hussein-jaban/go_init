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

	// Using the make function to create slices. 
	// NB: Slice always has underlying array, thats why they can have the same length but different capacity.

	h := make([]int, 3, 100)
	fmt.Printf("h: %v\n", h) // h: [0 0 0]
	fmt.Printf("Length: %v\n", len(h)) // Length: 3
	fmt.Printf("Capacity: %v\n", cap(h)) // Capacity: 100

	// splice concatination
	x := []int{1,2,3}
	x = append(x, []int{4,5,6}...) 
	fmt.Printf("x: %v\n", x) // x: [1 2 3 4 5 6]

	// removing element at the start
	y := x[1:]
	fmt.Printf("y: %v\n", y) // y: [2 3 4 5 6]
	
	// removing element at the end
	j := x[ : len(x) - 1]
	fmt.Printf("j: %v\n", j) // j: [1 2 3 4 5]
} 