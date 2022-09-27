package main

import "fmt"

/*
Anonymous functions
	• An anonymous function is a function that is declared without
	any named identifier to refer to it.

	• They can be used for containing functionality that need not be
	named and possibly for short term use.

	• They can accept inputs and return outputs, just as standard
	functions do.
*/

func main() {
	// assigning a fanuction declaration to a variable
	mult := func(first int, second int) int {
		return first * second
	}
	fmt.Printf("%T \n", mult) // type of mult: func(int, int) int
	fmt.Println(mult(10, 20)) // value: 200

	// call the function within the declaration and assign the result to a variable
	mult_var := func(first int, second int) int {
		return first * second
	}(10, 20) // inline call
	fmt.Printf("%T \n", mult_var) // type of mult_var: int
	fmt.Println(mult_var)         // value: 200
}
