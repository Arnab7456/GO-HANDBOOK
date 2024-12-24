package main

import "fmt"

func main() {
	var x int16 = 10
	var y int16 = 20
	var f float32 = 3.14
	var e float32 = 3.1444
	var Test bool = true
	var name string = "string"

	fmt.Println(x + y)
	fmt.Println(f+e)
	fmt.Println(x,f,Test, name);

	var test1 int16 = x + y;
	var test2  float32 = f+e;

	if test1 >int16(test2) {
		fmt.Println("test done")
	}else{
		fmt.Println("test failed")
	}
	for i:= 0; i<5;i++ {
		fmt.Println(i)
	}
	
}