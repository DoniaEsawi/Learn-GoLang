package main

import "fmt"

// global variables are declared here:
const pi float64 = 3.14
const sum = 3 // notice that the compiler here auto-assigned a data type <int> to the variable

var a string = "foo"

func main() {

	/*
		variables in inner blocks can access vars in outer blocks
		BUT NOT vice versa
		and ofcourse, global variables/consts can be accessed anywhere in the code
		Example:
	*/
	var b string = "bar"
	{
		var c string = "loreum"
		fmt.Println(b) // valid
		fmt.Println(c) // valid
		fmt.Println(a) //valid
	}

	// fmt.Print(c) // error
	fmt.Println(b)                        // valid
	fmt.Println(pi)                       // valid
	fmt.Printf("type of sum is: %T", sum) // valid

}
