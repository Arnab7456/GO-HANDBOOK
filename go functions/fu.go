package main

import "fmt"

func main() {
	getMessage()
	GreetHello("ans")
}

func getMessage() {
	fmt.Println("hello world")
}

func GreetHello(ans string) {
	fmt.Println(ans, "test ans")
}
