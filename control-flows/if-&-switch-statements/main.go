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

		// The concept of short-circuiting
		guess = -5
		if guess < 1 || returnTrue() || guess > 100 {
				fmt.Println("Your guess must be between 1 and 100")
		}

		// Checking multiple conditions using esle if 
		if false {
			fmt.Println("1st if condition")
		} else if true {
			fmt.Println("2nd if condition")
		} else {
			fmt.Println("3rd if condition")
		}

		// switch case statements
		switch 8 {
		case 1:
						fmt.Println("It is 1")
		// extending case test statements
		case 2, 4, 8:
						fmt.Println("It is are even numbers")
		default:
						fmt.Println("It is not 1 nor 2")
		}

		// Tagless switch case statements operations
		i := 5
		switch {
		case i > 1:
			fmt.Println("positive")
		case i < 1:
			fmt.Println("negative")
		default:
			fmt.Println("is one")
		}

}

func returnTrue() bool {
	fmt.Println("this function returns true")
	return true
}