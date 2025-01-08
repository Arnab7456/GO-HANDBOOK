package main

import "fmt"

// Generic function that works with any type
func PrintSlice[T any](slice []T) {
	for _, v := range slice {
		fmt.Println(v)
	}
}

func main() {
	nums := []int{1, 2, 3, 4}
	words := []string{"Go", "Generics", "Example"}

	// Use the generic function with integers
	PrintSlice(nums)

	// Use the generic function with strings
	PrintSlice(words)
}
