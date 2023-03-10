package main

import (
	"fmt"
)

func main()  {
		fmt.Println("If and switch statements init")

		// Simple if syntax
		if true {
				fmt.Println("This will be printed")
		}

		if false {
			fmt.Println("This will not be printed")
		}

		// The initializer syntax with if, ie pattern matching in elixir
		statePopulations := map[string]int{
				"California": 39250017,
				"Texas": 27862596,
				"Florida": 20612439,
				"New York": 19745289,
				"Pennsylvania": 12802503,
				"Illinois": 12801539,
				"Ohio": 11614373,
				}
		if pop, ok := statePopulations["Florida"]; ok {
					fmt.Println (pop)
		}
}