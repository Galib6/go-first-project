package main

import "fmt"

func addTwoNumber(a int, b int) int {
	return a + b
}

func main() {
	var sum = addTwoNumber(10, 20)
	fmt.Println(sum)
}
