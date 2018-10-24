package main

import (
	"fmt"
)

func main() {
	// //fix length array
	// var arr [3]int
	// arr[2] = 2
	// fmt.Println(arr)

	// //short hand fix array
	// arr := [3]int{1, 2, 3}
	// fmt.Println(arr)

	//short hand dynamic array
	arr := []int{1, 2, 3}
	arr = append(arr, 4)
	fmt.Println(arr)
}
