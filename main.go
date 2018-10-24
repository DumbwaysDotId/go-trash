package main

import (
	"fmt"
)

func main() {
	i := 10
	inc(&i)
	fmt.Println(i)
}

func inc(x *int) {
	*x++
}
