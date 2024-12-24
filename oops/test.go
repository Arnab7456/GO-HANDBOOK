package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{"test", 45}
	fmt.Println(p.Name, p.Age)
}