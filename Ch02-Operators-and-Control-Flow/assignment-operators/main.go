package main

import "fmt"

func main() {
	var x, y string = "foo", "bar"
	x += y
	fmt.Println(x) // "foobar"

	var z, s int = 27, 7
	z /= s
	fmt.Println(z) // 3

	var a, b float64 = 27.9, 7.0
	a -= b
	fmt.Println(a) // 20.9
	a += b
	fmt.Println(a) // 27.9

	var c, v int = 100, 9
	c /= v
	fmt.Println(c) // 11
	c %= v
	fmt.Println(c) // 2
}
