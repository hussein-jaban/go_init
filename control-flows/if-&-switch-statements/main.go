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
		// Providing true condition
		if pop, ok := statePopulations["Florida"]; ok {
					fmt.Println(pop)
		}
		// Providing false condition
		if pop, ok := statePopulations["bhbgrkjnfghirw"]; ok {
					fmt.Println(pop)
		}

		// using if with comparising and logical operator
		number := 50
		guess := 200
		if guess > number {
				fmt.Println("It is too hign")
		}
		if guess < number {
				fmt.Println("It is too low")
		}
		if guess == number {
				fmt.Println("Correct")
		}
		if guess < 1 || guess > 100 {
				fmt.Println("Your guess must be between 1 and 100")
		}

}