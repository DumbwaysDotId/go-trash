package main

import (
	"fmt"
	"time"
)

func main() {
	go timer("push up")
	go timer("sit up")

	fmt.Scanln()
}

func timer(todo string) {
	for i := 1; i < 5; i++ {
		fmt.Println(todo, i)
		time.Sleep(time.Second)
	}
}
