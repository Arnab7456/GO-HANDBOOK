package main

import "fmt"


func reverseArray(arr []int) {
	start := 0
	end := len(arr) - 1

	
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
}

// Helper function to print the array
func printArray(arr []int) {
	for _, val := range arr {
		fmt.Print(val, " ")
	}
	fmt.Println()
}

func main() {
	// Example array
	arr := []int{1, 2, 3, 4, 5}

	// Print the original array
	fmt.Println("Original Array:")
	printArray(arr)

	// Reverse the array
	reverseArray(arr)

	// Print the reversed array
	fmt.Println("Reversed Array:")
	printArray(arr)
}
