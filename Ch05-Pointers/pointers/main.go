package main

import "fmt"

/*
Pointers
	are variables that hold the memory address of another variables
	they also point to this address so that you can find or even change the values stored in that address

*/
func main() {

	// address operator: & -> typed before the variable and obtains the address of that variable
	// dereference operator: * -> typed before the address and returns the value stored in that address
	x := 5
	fmt.Println(&x)    // o/p: 0xc000018098
	fmt.Println(*(&x)) // o/p: 5

	fmt.Printf("%T \n", &x)    // o/p: *int
	fmt.Printf("%T \n", *(&x)) // o/p: int

	// declaring a pointer
	/*
		var <pointer_name> *<data_type>
	*/
	var ptr_i *int
	var ptr_s *string
	fmt.Println(ptr_i) // o/p: <nil> -> as it doesn't point to any address right now
	fmt.Println(ptr_s) // o/p: <nil>

	// initializing a pointer:

	/// method 1:
	// var <pointer_name> *<data_type>= &<variable_name>
	num := 10
	var ptr_num *int = &num
	fmt.Println(ptr_num)          // o/p: 0xc0000a60d0
	fmt.Println(*ptr_num)         // o/p: 10
	fmt.Printf("%T \n", ptr_num)  // o/p: *int
	fmt.Printf("%T \n", *ptr_num) // o/p: int

	/// method 2:
	// var <pointer_name>= &<variable_name> --> data type will be determined internally by the compiler
	num2 := 10
	var ptr_num2 = &num2
	fmt.Println(ptr_num2)          // o/p: 0xc000018140
	fmt.Println(*ptr_num2)         // o/p: 10
	fmt.Printf("%T \n", ptr_num2)  // o/p: *int
	fmt.Printf("%T \n", *ptr_num2) // o/p: int

	/// method 3:
	// <pointer_name> := &<variable_name> --> data type will be determined internally by the compiler
	num3 := 10
	ptr_num3 := &num3
	fmt.Println(ptr_num3)          // o/p: 0xc0000a6130
	fmt.Println(*ptr_num3)         // o/p: 10
	fmt.Printf("%T \n", ptr_num3)  // o/p: *int
	fmt.Printf("%T \n", *ptr_num3) // o/p: int

	// dereferencing a pointer: changing or accessing the value of the variable stored in the address that the pointer holds
	msg := "Unagi"
	fmt.Printf("%T %s\n", msg, msg)
	ptr_msg := &msg
	*ptr_msg = "Pivooot"
	fmt.Printf("%T %s\n", msg, msg)

	// pointers to an array
	arr := [3]int{1, 2, 3}
	ptr_arr := &arr
	fmt.Println(ptr_arr)          // o/p: &[1 2 3]
	fmt.Println(*ptr_arr)         // o/p: [1 2 3]
	fmt.Println(ptr_arr[0])       // o/p: 1
	fmt.Printf("%T \n", ptr_arr)  // o/p: *[3]int
	fmt.Printf("%T \n", *ptr_arr) // o/p: [3]int

}
