package main

import "fmt"

func main() {
	/*
	   switch expression {
	   case value_1:
	   // execute when expression equals to value_1
	   case value_2:
	   // execute when expression equals to value_2
	   default:
	   // execute when no match is found
	   }
	*/
	// example:
	day := "sunday"
	switch day {
	case "monday":
		fmt.Println("monday")
	case "tuesday":
		fmt.Println("tuesday")
	case "wednesday":
		fmt.Println("wednesday")
	case "thursday":
		fmt.Println("thursday")
	case "friday":
		fmt.Println("friday")
	case "saturday", "sunday": // we can add multiple value-conditions on the same case
		fmt.Println("weekend")
	default:
		fmt.Println("default")
	}
	// o/p: "weekend"
	/*
	   switch-condition case statement syntax:
	   switch {
	   case condition1:
	   // execute when condition1 is true
	   case condition2:
	   // execute when condition2 is true and condition1 is false
	   default:
	   // execute when no match is found
	   }
	*/
	// example:
	var a, b = 100, 5
	switch {
	case a/b == 10:
		fmt.Println("10")
	case a/b == 20:
		fmt.Println("20")
	case a/b >= 20:
		fmt.Println(">20")
	default:
		fmt.Println("default")
	}
	// o/p: 20

	// fallthrough -> a keyword is used in switch case to force the execution flow to fall through the successive case block.
	// example:
	day = "wednesday"
	switch day {
	case "monday":
		fmt.Println("monday")
	case "tuesday":
		fmt.Println("tuesday")
	case "wednesday":
		fmt.Println("wednesday")
		fallthrough
	case "thursday":
		fmt.Println("thursday")
		fallthrough
	case "friday":
		fmt.Println("friday")
		fallthrough
	case "saturday", "sunday":
		fmt.Println("weekend")
	default:
		fmt.Println("default")
	}
	//o/p:
	// wednesday
	// thursday
	// friday
	// weekend
}
