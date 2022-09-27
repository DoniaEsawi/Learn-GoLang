package main

import "fmt"

func main() {
	/*
	   syntax:
	   with default initialization:
	   var <array-name> [<size of the array>] <data type>

	   with manual initialization
	   var <array-name> [<size of the array>] <data type>=[<size of the array>] <data type> {data1, data2, ..}
	   or
	   <array-name>:=[<size of the array>] <data type> {data1, data2, ..}
	   or
	   <array-name>:=[...] <data type> {data1, data2, ..}

	*/
	// example 1:
	var arr [5]int
	fmt.Println(arr)
	// the above example is also eqivalent to:
	var arr0 [5]int = [5]int{}
	fmt.Println(arr0)
	// both outputs: [0 0 0 0 0]

	// example 2:
	arr1 := [2]string{"hello", "world"}
	fmt.Println(arr1)
	// this is eqivalent to:
	arr2 := [...]string{"hello", "world"}
	fmt.Println(arr2)
	// o/p: [hello world]

	// example 3:
	var arr3 [5]int = [5]int{1, 2, 3}
	fmt.Println(arr3)
	// o/p: [1 2 3 0 0]

	/* array indexing */

	arr[0] = 3
	fmt.Println(arr)
	// o/p: [3 0 0 0 0]
	fmt.Println(arr1[1])
	// o/p: world

	/* looping through array elements */

	// method 1:
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}
	// o/p:
	/*
		1
		2
		3
		0
		0
	*/

	// method 2:
	for index, element := range arr3 {
		fmt.Println("index", index, "=>", element)

	}
	// o/p:
	/*
		index 0 => 1
		index 1 => 2
		index 2 => 3
		index 3 => 0
		index 4 => 0
	*/

	// note: length = capacity in arrays
	fmt.Println("capacity of arr3 = ", cap(arr3), "\n",
		"length of arr3 = ", len(arr3))

	/* multi-dimensional arrays*/
	twoArr := [3][2]int{{1, 1}, {2, 3}, {4, 5}}
	fmt.Println(twoArr)       // o/p: [[1 1] [2 3] [4 5]]
	fmt.Println(twoArr[1][0]) // o/p: 2

}
