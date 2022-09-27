package main

import (
	"fmt"
)

/*
function declaration:
consists of 2 parts :

  - the function signature

  - the function body

    func <function_name>(<params>) <return type> {
    // the function body
    }
*/
const pi = 3.14

func sumFunc(op1 int, op2 int) int {
	// body
	return op1 + op2
}

// note: if we don't specify a return type, it'll assume the void return type
func printMsg(msg string) {
	fmt.Println("the message is:", msg)
}

// a function can return more than 1 value
func circle(rad float32) (float32, float32) {
	area := rad * rad * pi
	length := 2 * pi * rad
	return area, length
}

// you can declare the name of the returned parameter in the function signature
func square(side float32) (area float32, length float32) {
	area = side * side
	length = 4 * side
	return area, length
}

/*
Variadic functions
- functions that accepts variable number of arguments of the same type as the last parameter
func <func_name> (param1 type, param2 type, param3 ...type)<return type>

	{
		// body
	}

- here, param3 is actually a slice not a variable
*/
func sum(numbers ...int) int {
	sum := 0
	for _, val := range numbers {
		sum += val
	}
	return sum
}

func printGrades(studentName string, grades ...string) {
	fmt.Printf("%s grades are: ", studentName)
	for _, grade := range grades {
		fmt.Printf("| %s |", grade)
	}
	fmt.Print("\n")
}

func printDelayedMsg(msg string) {
	fmt.Println("Delayed Message: ", msg)
}
func main() {

	// calling a function
	sumOfNums := sumFunc(1, 2)
	fmt.Println(sumOfNums) // o/p: 3
	/////////////////////////////////
	/// defer satement: delays the execution of a function until the surrounding function returns.
	defer printDelayedMsg("I'm executed last")        // o/p:Delayed Message:  I'm executed last, this will be the last line of outputs
	defer printDelayedMsg("I'm executed before last") // notice that defer statement works in reverse order
	////////////////////////////////
	area, length := circle(5)
	fmt.Println("area of the circle = ", area)     // o/p: 78.5
	fmt.Println("length of the circle = ", length) // o/p: 31.400002
	////////////////////////////////
	sq_area, sq_length := square(5)
	fmt.Println("area of the square = ", sq_area)     // o/p: 25
	fmt.Println("length of the square = ", sq_length) // o/p: 20
	///////////////////////////////
	fmt.Println(sum(1, 2, 3, 4))                            // o/p: 10
	fmt.Println(sum())                                      // o/p: 0
	printGrades("Donia's", "A", "B+", "B", "A+", "C", "A+") // O/P: Donia's grades are: | A || B+ || B || A+ || C || A+ |
	//////////////////////////////
	// you can use 'blank identifier (_)' instead of declaring another variable that we don't need
	area_square, _ := square(5)
	fmt.Println(area_square) // o/p: 25

}
