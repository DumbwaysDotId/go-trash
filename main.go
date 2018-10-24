package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	// p := Person{Name: "Ega", Age: 28}

	// fmt.Println(p)

	// fmt.Println(p.Name)

	//array of struct
	p := []Person{
		{Name: "Ega", Age: 28},
		{Name: "Yuni", Age: 21},
	}

	// fmt.Println(p)
	// fmt.Println(p[1])
	// fmt.Println(p[1].Age)

	//traverse over array of struct
	for _, elem := range p {
		fmt.Println("Name", elem.Name)
		fmt.Println("Age", elem.Age)
		fmt.Println("")
	}
}
