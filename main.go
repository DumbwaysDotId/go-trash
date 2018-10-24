package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	result := sum(3, 2)
	fmt.Println(result)

	sqrtResult, err := sqrt(-2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sqrtResult)
	}
}

func sum(x int, y int) int {
	return x + y
}

func sqrt(num float64) (float64, error) {
	if num < 0 {
		return 0, errors.New("should not negative")
	}

	return math.Sqrt(num), nil
}
