package main

import "fmt"

// Define an interface
type Typer interface {
	Type() string
}

// Define a struct
type Keyboard struct{}

// Implement the Type method for Keyboard
func (k Keyboard) Type() string {
	return "Typing on a keyboard."
}

// Define another struct
type Smartphone struct{}

// Implement the Type method for Smartphone
func (s Smartphone) Type() string {
	return "Typing on a smartphone."
}

func main() {
	// Create instances of Keyboard and Smartphone
	var t Typer

	t = Keyboard{}
	fmt.Println(t.Type()) // Output: Typing on a keyboard.

	t = Smartphone{}
	fmt.Println(t.Type()) // Output: Typing on a smartphone.
}
