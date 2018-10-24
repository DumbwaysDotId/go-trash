package main

import (
	"fmt"
)

func main() {
	x := 2
	if x < 3 {
		fmt.Println("Ayee")
	} else if x == 4 {
		fmt.Println("is 4")
	} else {
		fmt.Println("Oops")
	}
}
