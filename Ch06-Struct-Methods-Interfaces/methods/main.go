package main

import "fmt"

type Square struct {
	side_length int
	area        int
}

/*
Methods:
	are functions that has a defined (reciever), a reciever is typically a struct,
	so the methods are functions associated with a specific data type or struct,
	 only entities of that struct can call them
*/

// example 1: Method that can modify the values of the original struct object
func (sq *Square) calcArea() {
	sq.area = sq.side_length * sq.side_length
}

// example 2: Method that can't modify the values of the original struct object
func (sq Square) calcAreaValue() {
	sq.area = sq.side_length * sq.side_length
}

func main() {
	// test example 1
	square := Square{
		side_length: 5,
	}
	square.calcArea()
	fmt.Println(square.area) // o/p: 25

	// test example 2
	square2 := Square{
		side_length: 4,
	}
	square2.calcAreaValue()
	fmt.Println(square2.area) // o/p: 0

}
