package main

import (
	"fmt"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog,
	}
	fmt.Println(fido)
	fmt.Println(dog.YearsTwo(20))
}
