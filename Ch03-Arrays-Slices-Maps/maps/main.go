package main

import (
	"fmt"
)

func main() {
	/*
		maps is a way to store data in a key-value pair manner
	*/

	// initialization
	my_map := map[string]int{"age": 10, "price": 200}
	fmt.Println(my_map) // o/p: map[age:10 price:200]

	var temp_map map[string]int // creates a nil map
	//temp_map["id"] = 10 -> error, assignment to entry in nil map
	fmt.Println(temp_map) // o/p: map[]

	map_mk := make(map[int]string) // creates a nil map
	map_mk[1] = "good"             // valid
	fmt.Println(map_mk)            // o/p: map[1:good]

	// adding a key-value pair
	spec := map[string]string{
		"color":   "red",
		"price":   "200",
		"brand":   "MG",
		"country": "UAE",
	}
	spec["year"] = "2022"
	fmt.Println(spec) // o/p: map[brand:MG color:red country:UAE price:200 year:2022]
	value, found := spec["name"]
	fmt.Println("name key found:", found, "and its value: ", value)
	// updating a key-value pair

	spec["price"] = "5000"
	fmt.Println(spec) // o/p: map[brand:MG color:red country:UAE price:5000 year:2022]

	// deleting a key-value pair
	delete(spec, "price")
	fmt.Println(spec) // o/p: map[brand:MG color:red country:UAE year:2022]

	// iterating over a map
	for key, value := range spec {
		fmt.Println(key, "=>", value)
	}
	/*
		o/p:
			color => red
			brand => MG
			country => UAE
			year => 2022
	*/

	// clearing a map
	for key, _ := range spec {
		delete(spec, key)
	}
	fmt.Println(spec) // o/p: map[]

}
