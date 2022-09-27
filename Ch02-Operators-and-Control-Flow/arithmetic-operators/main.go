package main

import "fmt"

func main() {

	// addition (+)
	// integers
	fmt.Println(3 + 5) // o/p: 8
	// floats
	fmt.Println(3.2 + 5) // o/p: 8.2

	// strings
	fmt.Println("Hello" + "World") // o/p: HelloWorld
	// note: only addition is allowed in string variables

	// int , float
	/*
		var a int = 2
		fmt.Println(a + 5.51)
		// o/p: error: 5.51 (untyped float constant) truncated to int
	*/

	var a float64 = 5.51
	fmt.Println(a + 2) // o/p: 7.51

	// int , string
	/*
		fmt.Println(a + "cats")
		// o/p: error: invalid operation: a + "cats" (mismatched types int and untyped string)
	*/

	// subtraction (-)
	fmt.Println(3 - 5) // o/p: -2
	// fmt.Println("hi" - "there") -> invalid

	// multiplication (*)
	fmt.Println(3 * 5) // o/p: 15
	// fmt.Println("hi" * "there") -> invalid

	// division (/)
	fmt.Println(3 / 5) // o/p: 0

	// modulus (%)
	fmt.Println(3 % 5) // o/p: 3

	/* Unary Operators */
	// increment (++)
	b := 1
	b++
	fmt.Println(b) // o/p: 2

	// decrement (--)
	b--
	fmt.Println(b) // o/p: 1
}
