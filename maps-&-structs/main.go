package main

import "fmt"

func main()  {
	// Creating Map
	statePopulation := map[string]int{
		"ohio": 345656,
		"columbus": 36457537,
		"new york": 98798688,
	}
	fmt.Println(statePopulation) // map[columbus:36457537 new york:98798688 ohio:345656]

	// Creating map with the make function
	make_map := make(map[string]int)
	make_map = map[string]int{
			"a": 1,
			"b": 2,
			"c": 3,
	}
	fmt.Println(make_map) // map[a:1 b:2 c:3]

	// Accessing values from a map
	fmt.Println(statePopulation["ohio"]) // 345656
	fmt.Println(make_map["c"]) // 3

	// Adding key value pairs to an existing map
			// The oeder of a map is not guaranteed unlike arrays and slices.
	fmt.Println(make_map) // map[a:1 b:2 c:3]
	make_map["d"] = 4
	fmt.Println(make_map) // map[a:1 b:2 c:3 d:4]

	// Deleting
	delete(make_map, "b")
	fmt.Println(make_map) // map[a:1 c:3 d:4]


	// Using coma okay syntax to check if a key exists and returns true if it does or false if it doesn't

	a, ok := make_map["a"]
	fmt.Println(a, ok) // 1 false

	z, ok := make_map["z"]
	fmt.Println(z, ok) // 0 false




}