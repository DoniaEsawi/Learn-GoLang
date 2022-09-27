package main

import "fmt"

/*
high order function
	function that receives a function as an argument or returns a
	function as output.
why we use it?
	• creating smaller functions that take care of certain piece of logic.
	• composing complex function by using different smaller functions.
*/
const pi = 3.14

func getArea(rad float32) float32 {
	return rad * rad * pi
}

func getLength(rad float32) float32 {
	return 2 * rad * pi

}

func getDiameter(rad float32) float32 {
	return 2 * rad

}

// note here the function takes a function as an input
func printResult(radius float32, calcFunction func(r float32) float32) {
	result := calcFunction(radius)
	fmt.Println("Result: ", result)
}

// note here the return type a this function is a function too
func getFunction(query int) func(r float32) float32 {

	queryOfFunc := map[int]func(r float32) float32{
		1: getDiameter,
		2: getLength,
		3: getArea,
	}
	return queryOfFunc[query]

}

func main() {
	var query int
	var radius float32
	fmt.Print("Enter the radius of the circle:")
	fmt.Scanf("%f ", &radius)
	fmt.Printf("Choose what to calculate: \n 1- diameter \n 2- length \n 3- area:\n")
	fmt.Scanf("%d ", &query)

	printResult(radius, getFunction(query))
}
