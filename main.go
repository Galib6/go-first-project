package main

import "fmt"

func call() {
	var x []int

	x = append(x, 10)
	x = append(x, 20)
	x = append(x, 30)

	y := x

	y = append(y, 20)
	y = append(y, 60)
	y = append(y, 90)

	x[0] = 3

	fmt.Println(x, cap(x))
	fmt.Println(y, cap(y))
}

func main() {
	call()
}
