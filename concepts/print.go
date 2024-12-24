package main

import "fmt"

func main() {
	const name, age = "Kim", 22
	//
	fmt.Print(name, " is ", age, " years old.\n")  // Kim is 22 years old.
	fmt.Printf("%v is %v years old.\n", name, age) // Kim is 22 years old.
	fmt.Println(name, "is", age, "years old.")     // Kim is 22 years old.
	//
	print(name, " is ", age, " years old.\n") // Kim is 22 years old.
	println(name, "is", age, "years old.")    // Kim is 22 years old.
}
