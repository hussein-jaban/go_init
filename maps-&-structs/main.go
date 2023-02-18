package main

import (
	"fmt"
	"reflect"
)

// Structs
type Doctor struct {
			number int
			actorName string
			companions []string
}
type Animal struct {
	// using tags for validations
	name string `required max:"100"`
	origin string
}
// embeded structs like inheritance in OOP
type Bird struct {
		Animal
		speedKPH int
		canFly bool
}

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
	fmt.Println(a, ok) // 1 true

	z, ok := make_map["z"]
	fmt.Println(z, ok) // 0 false

	// impelenting structs
	aDoctor := Doctor {
			number: 3,
			actorName: "Jon Pertwee",
			companions: [] string {
					"Jo Liz Sarah", "Grant Shaw",  "Jane Smith",
			},
	}
 fmt.Println(aDoctor) // {3 Jon Pertwee [Jo Liz Sarah Grant Shaw Jane Smith]}

	// Inspecting structs with the dot notation
 fmt.Println(aDoctor.number) // 3
 fmt.Println(aDoctor.actorName) // Jon Pertwee
 fmt.Println(aDoctor.companions) // [Jo Liz Sarah Grant Shaw Jane Smith]

	// Declaring structs anonymosly ie with out the type keyword

	user := struct{name string; age int} {name: "hussein", age: 22}
	fmt.Println(user) // {hussein 22}

	// Embedded structs
	bird0 := Bird{
		Animal: Animal{name: "chicken", origin: "usa"},
		speedKPH: 20,
		canFly: false,
	}
	fmt.Println(bird0) // {{chicken usa} 20 false}


	bird1 := Bird{}
	bird1.name = "Ostrich"
	bird1.origin =  "Africa"
	bird1.speedKPH=  45
	bird1.canFly=  false

	fmt.Println(bird1) // {{Ostrich Africa} 45 false}

	// Validation with tags in structs
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("name")
	fmt.Println(field.Tag) // required max:"100"

}