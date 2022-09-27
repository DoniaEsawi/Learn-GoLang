package main

import "fmt"

func main() {
	a := true
	b := false
	c := true
	d := false
	/* AND operator && */
	fmt.Println(a && b) // false
	fmt.Println(a && c) // true
	fmt.Println(b && d) // false
	/* OR operator || */
	fmt.Println(a || b) // true
	fmt.Println(a || c) // true
	fmt.Println(b || d) // false
	/* NOT operator ! */
	fmt.Println(!b) // true
	fmt.Println(!a) // false

	// note: Logical operators ONLY work with boolean data types
	/*
		num := 3
		fmt.Println(num && 0)
	*/
}
