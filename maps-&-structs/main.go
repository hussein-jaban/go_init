package main

import "fmt"

func main()  {
	statePopulation := map[string]int{
		"ohio": 345656,
		"columbus": 36457537,
		"new york": 98798688,
	}
	fmt.Println(statePopulation) // map[columbus:36457537 new york:98798688 ohio:345656]
}