package main

import "fmt"

func main() {
	var number [6]int

	number[0] = 4
	number[1] = 3
	number[2] = 2
	number[3] = 1
	number[4] = 0
	fmt.Println(number)
	fmt.Println(AppendTest())	
}

func AppendTest ()[] int{
	var number [] int 

	number = append(number, 5,4,3,2,1,0);
	return number
}