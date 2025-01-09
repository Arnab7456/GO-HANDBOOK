package main

import "fmt"
import "package/ans"
func main() {
	data := []float64{4, 5, 10, 15}
	fmt.Println(data)
	avg := ans.avg(data)
	fmt.Println(avg)
}