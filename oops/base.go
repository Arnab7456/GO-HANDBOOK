package main

import "fmt"


type Sum struct{}

func (s Sum) SumInt2(x, y int16) int16 {
	return x + y
}

func (s Sum) SumInt3(x, y, z int) int {
	return x + y + z
}


func (s Sum) SumFloat2(x, y float64) float64 {
	return x + y
}

func main() {
	s := Sum{}

	
	fmt.Println(s.SumInt2(10, 20))       
	fmt.Println(s.SumInt3(10, 20, 30))   
	fmt.Println(s.SumFloat2(10.5, 20.5)) 
}
