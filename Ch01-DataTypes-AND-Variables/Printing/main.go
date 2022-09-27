package main

import "fmt"

/* Types of print functions in Go */

func main() {
	// fmt.Print(msg1, msg2, ..) -> prints a string or any other data type without an endline character at the end
	fmt.Print("There are ", 22, "Strudents in the class who said it was ", true, "\n") // o/p: HelloThere22 true
	// fmt.Println ->	//prints a string or any other data type with an endline character at the end
	fmt.Println("There are ", 22, "Strudents in the class who said it was ", true) // o/p: HelloThere22 true
	// fmt.Printf -> format specifier - prints the variables in a specific format
	/*
		%v -> default format
		%T -> type of the value
		%d -> integers
		%c -> character
		%q -> quoted character / string
		%s -> plain string
		%t -> true or false
		%f -> floating numbers
		%0.2f -> floating numbers up to 2 decimal places
	*/
	var score float64 = 32.556
	var name string = "Regina Phalange"
	var isLiked bool = true
	var fullScore int = 100
	fmt.Printf("%q's score in sarcasm is %0.1f out of %d, but do we still like her? \n answer is: %t",
		name, score, fullScore, isLiked)
}
