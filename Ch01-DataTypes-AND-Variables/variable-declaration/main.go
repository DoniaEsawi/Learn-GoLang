package main

import "fmt"

func main() {

	var name string = "Regina Phalange"
	var (
		name2 string = "Princess Concuella"
		age   int    = 21
	)
	var joeyName, reginaName string = "kenAdams", "Pheobe"
	fmt.Println(name, name2, age, joeyName, reginaName)

	donia := "psycho"
	// donia = 1 --> error!
	fmt.Println(donia)

	/*

		error:
		var s,t string,int = "Priyanka", 100
		- multiple variable declaration must be of the same type

	*/

}
