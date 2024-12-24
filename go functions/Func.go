package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	res := add(10, 55)
	fmt.Println(res)
}