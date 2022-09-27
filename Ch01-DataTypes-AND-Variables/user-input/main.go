package main

import "fmt"

func main() {

	var firstName string
	var age int
	var isEnrolled bool
	/*
		The general format of Scanln is
		fmt.Scanln(&var1Name, &var2Name, ...)
	*/

	// example 1:
	fmt.Println("Enter your first Name")
	fmt.Scanln(&firstName)
	fmt.Println("your entered name: ", firstName)

	/*
		o/p 1:
		Enter your first Name
		> donia
		your entered name:  donia

	*/

	// example 2:
	fmt.Println("Enter your age and if enrolled: Age Enrolled?")
	fmt.Scanln(&age, &isEnrolled)
	fmt.Println("your entered age: ", age)
	fmt.Println("your entered boolean: ", isEnrolled)

	/*
		o/p 2:
		Enter your age and if enrolled: Age Enrolled?
		> 21 true
		your entered age:  21
		your entered boolean:  true
	*/

	////////////////////////////////////////////////////////////

	/*

		The general format of Scanf is
		fmt.Scanf("%<format specifier>", &varName)

		important note: if you use scanf, it will store the Enter
		key entered by the user in the input buffer,
		so if you want to use it, make it the last scan in the code,
		or make all the scans in the same scanf function call

		notice that Scanf returns 2 variables:
		- count: which displays the count of inputs
		- error: to indicate whether an error occurred or not
	*/

	//example 3:
	var temp int
	fmt.Println("Enter your age as a string: ")
	count, err := fmt.Scanf("%d", &temp)
	fmt.Println("Entered age is: ", temp)
	fmt.Println("count is: ", count)
	fmt.Println("err is: ", err)
	fmt.Printf("type of count var is: %T\n", count) // -> int
	fmt.Printf("type of err var is: %T\n", err)     // -> *errors.errorString

	/*
		o/p 3:
		Enter your age as a string:
		> hi
		Entered age is:  0
		count is:  0
		err is:  expected integer
		type of count var is: int
		type of err var is: *errors.errorString

	*/

}
