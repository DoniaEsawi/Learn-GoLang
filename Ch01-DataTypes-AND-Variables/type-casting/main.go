package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	/*
		first: lets get the tpye of the vriables
		there are 2 ways for that:
		1- using reflect.TypeOf(var)
		2- using the %T formatter
	*/
	msg := "hello"
	pi := 3.14

	fmt.Println(reflect.TypeOf(msg)) // string
	fmt.Printf("%T\n", pi)           // float64

	/*
		now let's try to convert types
	*/
	// int to float
	var percent int = 80
	fmt.Println(float64(percent))
	// float to int
	fmt.Println(int(pi))
	/*
		string-int conversions
		note that Atoi returns 2 outputs:
		1st o/p: is the result of convertion
		2nd o/p is the error
	*/
	// int to string
	fmt.Println(strconv.Itoa(percent))
	// string to int
	out, err := strconv.Atoi(msg)
	fmt.Println("o/p:", out)
	fmt.Println("error:", err)

	out2, err2 := strconv.Atoi("200")
	fmt.Println("o/p:", out2)
	fmt.Println("error:", err2)
	/*
		final o/p
		string
		float64
		80
		3
		80
		o/p: 0
		error: strconv.Atoi: parsing "hello": invalid syntax
	*/
}
