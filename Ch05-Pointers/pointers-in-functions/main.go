package main

import "fmt"

/*
Passing by Value in functions
the parameter is copied to another location in the memory,
so when we pass an argument to the function and modify it inside the function, only the copied value is modified,
and the original value will not change
*/
// e.g.
func modify_variable(value int) {
	value += 10
}

/*
Passing by reference in functions
the parameter address is passed into the function call,
and hence the original value is changed, and the value isn't copied to another location
*/
// e.g.
func change_variable_value(value *int) {
	*value += 10
}

func change_slice_value(slice []int) {
	slice[0] += 10
}

func change_map_value(m map[string]string) {
	_, found := m["new"]
	if !found {
		m["new"] = "modified"
	}
}
func main() {
	// testing passing by value
	val := 20
	fmt.Println(val) // o/p: 20
	modify_variable(val)
	fmt.Println(val) // o/p: 20 --> note that the original var wasn't modified
	// testing passing by value
	val2 := 20
	fmt.Println(val2)            // o/p: 20
	change_variable_value(&val2) // note that we pass the address of the variable not a copy of it
	fmt.Println(val2)            // o/p: 30 --> note that the original var was modified

	// recall: a slice itself is a pointer (reference) to an array
	// so passing a slice to a function is like passing a pointer, the original values will be modified
	// e.g.
	slice := []int{1, 2, 3}
	fmt.Println(slice) // o/p: [1 2 3]
	change_slice_value(slice)
	fmt.Println(slice) // o/p: [11 2 3]

	// Maps: maps also are passed by reference by default
	my_map := map[string]string{
		"color":   "red",
		"price":   "200",
		"brand":   "MG",
		"country": "UAE",
	}
	fmt.Println(my_map) // o/p: map[brand:MG color:red country:UAE price:200]
	change_map_value(my_map)
	fmt.Println(my_map) // o/p: map[brand:MG color:red country:UAE new:modified price:200]
	// notice how the new key is added to the map
}
