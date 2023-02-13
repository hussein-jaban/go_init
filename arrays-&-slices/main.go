package main

import (
	"fmt"
)

func main()  {
	// Arrays
	grades :=[3]int{97, 85, 93}
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
} 