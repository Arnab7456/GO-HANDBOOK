package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y)) // 5 => 9 +16
	fmt.Println(x, y, f)
	test();
}

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func test() {
	// like for in range loop
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

/*
 like py
for x in range(6):
  print(x)
*/