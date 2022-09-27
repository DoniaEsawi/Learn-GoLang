package main

import "fmt"

/*
Structs:
	are user-defined data type that groups together multiple values or data elements
	and provides a way to reference them through a single variable name (entity)

Syntax:
	type <struct name> struct{
		// field 1
		// field 2
		.
		.
	}

*/
// example
type Car struct {
	model  string
	car_no string
	color  string
	id     int
}
type Driver struct {
	name string
	age  int
	id   int
}

func increase_car_id_reference(car *Car) {
	car.id++
}
func increase_car_id_value(car Car) {
	car.id++
}

func main() {
	// initialization
	/*
		var <variable_name> <struct name>
	*/
	// method 1
	var car Car
	// "%+v" prints the data names and the values of a struct entity
	fmt.Printf("%+v\n", car) // o/p: {model: car_no: color: id:0}

	// method 2
	/*
		<variable_name>:= new(<struct_name>)
		the 'new' keyword, allocates new memory addresses to every field
		in the struct and returns a pointer to that struct
		*not a very common way*
	*/
	car_ptr := new(Car)
	fmt.Printf("%+v\n", car_ptr)  // o/p: &{model: car_no: color: id:0}
	fmt.Printf("%+v\n", *car_ptr) // o/p: {model: car_no: color: id:0}
	fmt.Printf("%T\n", car_ptr)   // o/p: *main.Car

	// method 3 (the most common way)
	/*
		<variable_name>:= <struct_name> {
			<filed_name> : <value>,
			<filed_name> : <value>,
			<filed_name> : <value>,
		}
	*/
	my_car := Car{
		id:     1,
		model:  "MG5",
		color:  "Red",
		car_no: "1234",
	}
	fmt.Printf("%+v\n", my_car) // o/p: {model:MG5 car_no:1234 color:Red id:1}

	// can also assign the values of a struct in their order without specifying the field names (not a good practice)
	my_car2 := Car{
		"654",
		"Kia-Sprotage",
		"Black",
		2,
	}
	fmt.Printf("%+v\n", my_car2) // o/p: {model:MG5 car_no:1234 color:Red id:1}

	// accessing a field
	my_car2.id = 3
	my_car2.color = "Silver"
	// my_car2.price = 200 --> error (unkown field name)
	fmt.Printf("%+v\n", my_car2) // o/p: {model:654 car_no:Kia-Sprotage color:Silver id:3}

	/* Passing Structs to a function */

	// by value
	increase_car_id_value(my_car2)
	fmt.Printf("%+v\n", my_car2) // o/p: {model:654 car_no:Kia-Sprotage color:Silver id:3}
	// the id is not increased

	// by reference
	increase_car_id_reference(&my_car2)
	fmt.Printf("%+v\n", my_car2) // o/p: {model:654 car_no:Kia-Sprotage color:Silver id:4}
	// the id is increased

	/* Comparing Structs */

	// structs of the same type can be compared using the Go'd equality operators "==" or "!="
	isEqual := my_car2 == my_car
	fmt.Printf("%t\n", isEqual) // o/p: false
	my_car2 = my_car
	isEqual = my_car2 == my_car
	fmt.Printf("%t\n", isEqual) // o/p: true

	driver := Driver{
		name: "Ali",
		age:  35,
		id:   1,
	}
	//isEqual = my_car == driver // error: invalid operation: my_car == driver (mismatched types Car and Driver)
	fmt.Println(driver) // o/p: {Ali 35 1}

}
