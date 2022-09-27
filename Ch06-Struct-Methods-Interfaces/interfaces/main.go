package main

import "fmt"

/*
Interface:
	a powerful way to introduce modularity in GoLang
	they specify a set of methods but don't implement them

syntax:
type <interface-name> interface{
	// Method signatures

}
a single interface can has any number of methods but only 1 method set
method set: the set of methods that has the same reciever (struct)
*/
// example:
type shape interface {
	area() float64
	perimeter() float64
}

/*
Implementing an interface
	a type implements an interface by implementing its methods
	the goLang interfaces are implemented implicitly and there's no specific keyword to implement them
*/

type Square struct {
	side_length float64
}

func (square Square) area() float64 {
	return square.side_length * square.side_length
}
func (square Square) perimeter() float64 {
	return square.side_length * 4.0
}

type Rectangle struct {
	width  float64
	height float64
}

func (rect Rectangle) area() float64 {
	return rect.width * rect.height
}
func (rect Rectangle) perimeter() float64 {
	return 2 * (rect.height + rect.width)
}

// note that the 2 struct types, implements all the methods of the interface shape,
// so they implement the interface shape
// note: to say that a struct implements an interface, it must implements all the methods of that interface

func printS(s shape) {
	fmt.Println("dimensions: ", s)
	fmt.Println("area: ", s.area())
	fmt.Println("perimeter: ", s.perimeter())
}

func main() {
	rect := Rectangle{
		width:  5,
		height: 4,
	}
	square := Square{
		side_length: 5,
	}

	fmt.Println("Rectangle:")
	printS(rect)
	/*
		dimensions:  {5 4}
		area:  20
		perimeter:  18
	*/

	fmt.Println("Square:")
	printS(square)
	/*
		dimensions:  {5}
		area:  25
		perimeter:  20
	*/

	// so, we were able to print the details of 2 different structs using the same function

}
