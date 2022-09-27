package main

import "fmt"

func main() {
	/*
	   general syntax for for-loops

	   for intialization; condition; post{
	   	// do something
	   }
	   ● both intialization and post parts are optional
	   ● if you removed the condition part it will become an infinite loop
	*/
	// example 1:
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// example 2:
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// example 3 (infinite loop):
	/*for {
		fmt.Println(i)
	}*/

	// o/p
	/*
		0
		1
		2
		3
		4
	*/

	/* continue & break */
	for i := 0; i < 9; i++ {
		if i%2 == 1 {
			continue
		}
		if i == 8 {
			break
		}
		fmt.Println(i) // prints only even numbers < 8
	}
	// o/p
	/*
		0
		2
		4
		6
	*/
}
