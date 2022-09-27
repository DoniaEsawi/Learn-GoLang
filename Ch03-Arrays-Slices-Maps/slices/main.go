package main

import "fmt"

func main() {
	/*
		Slices are pointers to array elements
			- so when you change the value of element in the slice, this change will be reflected in the array too
		Slices has a capacity = len(array) - starting index of the slicing
		Unlike Arrays, Slices are of dynamic size and are more flexible
	*/

	// initialization
	students := []string{"adam", "sara", "emily"}
	fmt.Println(students) // o/p: [adam sara emily]

	// using make function
	slice_mk := make([]int, 3, 5) // type, length, capacity
	fmt.Println(slice_mk)         // o/p: [0 0 0]
	fmt.Println(len(slice_mk))    // o/p: 3
	fmt.Println(cap(slice_mk))    // o/p: 5
	// pointing to an array

	arr := [5]int{2, 3, 4, 5, 6}
	slice := arr[1:4]       // note that the first index is included int the slice but the last index is not included
	fmt.Println(slice)      // o/p: [3 4 5]
	fmt.Println(len(slice)) // o/p: 3
	fmt.Println(cap(slice)) // o/p: 4

	slice[0] = 88
	fmt.Println(arr) // o/p: [2 88 4 5 6] -> notice that the change in slice was reflected in the array it points to

	// appending to a slice

	// appending an element to a slice
	slice = append(slice, 10)
	fmt.Println(slice)      // o/p: [88 4 5 10]
	fmt.Println(cap(slice)) // o/p: 4

	slice = append(slice, 20, 30, 40)
	fmt.Println(slice)      // o/p: [88 4 5 10 20 30 40]
	fmt.Println(cap(slice)) // o/p: 8 --> the capacity of the slice is got doubled as the old capacity was < the length of the slice, i.e. the original slice is appended with a duplicate slice

	// appending a slice to another slice

	full_arr := [6]int{2, 4, 6, 3, 5, 7}
	even_slice := full_arr[0:3]
	odd_slice := full_arr[3:5]

	full_slice := append(even_slice, odd_slice...)

	fmt.Println(full_slice)      // o/p: [2 4 6 3 5]
	fmt.Println(cap(full_slice)) // 6

	// deleting from a slice
	// create 2 slices (one ends at the index we want to delete, and the other starts from that index +1) and append them together

	arr2 := [5]int{1, 2, 0, 4, 5}
	fmt.Println(arr2) // o/p: [1 2 0 4 5]
	slice1 := arr2[0:2]
	slice2 := arr2[3:5]
	new_slice := append(slice1, slice2...)

	fmt.Println(new_slice)      // o/p: [1 2 4 5]
	fmt.Println(cap(new_slice)) // o/p: 5

	fmt.Println(arr2) // o/p: [1 2 4 5 5] -> note that it's deleted from the array but since the array size is static, it 'll duplicate the last element to fill the length

	// copying from a slice

	src_slice := []int{1, 2, 3, 4}
	dest_slice := make([]int, 3)

	res := copy(dest_slice, src_slice) // returns the number of elements that's copied which is = min(len(dest, src))
	fmt.Println(res)                   // o/p: 3

}
