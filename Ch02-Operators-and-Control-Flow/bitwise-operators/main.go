package main

import "fmt"

func main() {
	var x, y int = 100, 90
	/*
		x = 0110 0100
		y = 0101 1010
	*/
	// bitwise-AND
	fmt.Println(x & y) // 0100 0000 = 64
	// bitwise-OR
	fmt.Println(x | y) // 0111 1110 = 126
	// bitwise-XOR
	fmt.Println(x ^ y) // 0011 1110 = 62
	// left shift operator
	// x + y = 190 = 1011 1110
	fmt.Println((x + y) << 2) // 0010 1111 1000 = 760
	// right shift operator
	// x + y = 190 = 1011 1110
	fmt.Println((x + y) >> 2) // 0010 1111 = 47

}
