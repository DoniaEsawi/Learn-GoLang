package main

import "fmt"

func main() {
	/*
	   if
	   condition_1 {

	   execute when condition_1 is true

	   } else if condition_2 {

	   execute when condition_1 is false,
	   and condition_2 is true

	   } else if condition_3 {

	   execute when condition_1 and 2 are false,
	   and condition_3 is true

	   } else {
	    when none of the above conditions are true
	   }
	*/

	// example:
	var a, b string = "foo", "bar"
	if a+b == "foo" {
		fmt.Println("foo")
	} else if a+b == "foobar" {
		fmt.Println("bar")
	} else if a+b == "foobar" {
		fmt.Println("foobar")
	} else {
		fmt.Println("None matched")
	}
	fmt.Println("thank you!")

	// o/p:
	// bar
	// thank you!

	// notes: you can use parentheses () around the condition but it's optional, and actually
	// if you enabled the auto-formatting option, it will be automatically deleted :)
	if 3 < 5 {
		fmt.Println("correct!")
	}
	// important note: else {} must begin IN THE SAME LINE that the if block ended,
	// also the starting curling brace { must be in the same line of the if or else
	/*
		if(3>5)
		{

		}
		// invalid

		if(3>5) {

		}
		else {

		}
		// invalid
	*/

}
