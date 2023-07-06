package main

import (
	"fmt"
	"log"
)

func main() {
	//call sum function and print the result
	fmt.Println(sum(5, 5))
	fmt.Println("Hello World")
	log.Println("Hello World")
}

// func sum to add two numbers
func sum(x int, y int) int {
	return x + y
}
