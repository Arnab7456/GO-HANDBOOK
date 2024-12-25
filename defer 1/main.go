package main

import "fmt"

func add(a, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a - b
}

func multi(a, b int) int {
	return a * b
}


func main() {
	fmt.Println(add(5, 7))
	data := add(5, 5)
	data1 := sub(10,5);
	data2 := multi(10,10);
	fmt.Println(data)
	defer fmt.Println(data1)
	fmt.Println(data2)
	fmt.Println("testing golang 1")
	defer fmt.Println("testing golang 2")
	fmt.Println("testing golang 3")
}
