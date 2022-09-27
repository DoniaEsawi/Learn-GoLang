package main

import "fmt"

/*
Recursive functions
	The function that keep calling itself untill it reaches a base case
*/
func factorial(n int) int {
	if n < 0 { // corner case
		return -1 // invalid
	}
	if n == 0 || n == 1 { // base case
		return 1
	}
	return n * factorial(n-1) // recursive call

}
func main() {
	fmt.Println(factorial(4))  // o/p: 24
	fmt.Println(factorial(0))  // o/p: 1 -> base case
	fmt.Println(factorial(1))  // o/p: 1 -> base case
	fmt.Println(factorial(-5)) // o/p: -1 -> corner case

}
