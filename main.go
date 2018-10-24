package main

import (
	"fmt"
)

func main() {
	// //standard forloop
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// }

	// //while loop
	// i := 0
	// for i < 5 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// //array range
	// arr := []int{1, 2}

	// for index, value := range arr {
	// 	fmt.Println("index", index, "value", value)
	// }

	//map loop
	bio := make(map[string]string)

	bio["firstName"] = "Ega"
	bio["lastName"] = "Radiegtya"

	for key, value := range bio {
		fmt.Println("key", key, "value", value)
	}
}
