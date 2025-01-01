package main

import (
	"fmt"
	"math"
)

type Datax struct {
	x, y float64
}

func (v Datax) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

//  a method on non-struct types, too.
type TestFloat float64

func (f TestFloat) testAbs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	v := Datax{4, 2}
	f := TestFloat(-math.Sqrt2)
	fmt.Println(v.Abs())
	fmt.Println(f.testAbs())
}
