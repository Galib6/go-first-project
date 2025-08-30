package main

import "fmt"

var (
	a = 10
)

func main() {
	fmt.Println("Hello Init function")
}

func init() {
	fmt.Println("I am the first function that executes first")
}
