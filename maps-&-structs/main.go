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

}