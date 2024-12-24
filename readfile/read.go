package main

import (
	"fmt"
	"os"
)

func main() {
	res , err := os.ReadFile("data.txt");
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(res))
}	