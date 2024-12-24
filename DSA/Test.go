package main

import "fmt"

func search(arr []int, x int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == x {
			return i
		}
	}
	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	x := 4

	res := search(arr, x)
	if res == -1 {
		fmt.Println("Element is not found in the array")
	} else {
		fmt.Println("Element is found", res)
	}

}
