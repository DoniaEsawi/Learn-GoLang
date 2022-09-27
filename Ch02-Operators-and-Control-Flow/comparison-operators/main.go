package main

import "fmt"

func main() {
	a := 1
	b := 1
	// ==
	fmt.Println(a == b) // true
	// !=
	b = 2
	fmt.Println(a != b) //true
	// >=
	fmt.Println(a >= b) //false
	// <=
	fmt.Println(a <= b) //true
	// >
	fmt.Println(a > b) //false
	// <
	fmt.Println(a < b) //true

	// note that you cannot compare 2 variables with different data types
	/*
		var a string = "cat"
		var b int = 12
		fmt.Println(a == b)
	*/
	// error:  invalid operation: a == b (mismatched types string and int)

}
