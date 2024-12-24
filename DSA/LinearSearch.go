package main

import (
	"fmt"
	"sort"
)

func main() {
	res := []int{1, 2, 4, 5, 9, 8, 7}
	var target int = 10
	found := false
	sort.Ints(res)
	for i := 0; i < len(res); i++ {
		if res[i] == target {
			fmt.Println("target found", res[i])
			found = true
			break
		}
	}
	if !found {
		fmt.Println("target not found")
	}

	fmt.Println(res)

}
