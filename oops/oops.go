package main

type Area struct {}

func (s Area) Area(l,b int) int{
	return l*b
}

func main (){
	// java like calling
	area := Area{}
	println(area.Area(10,20))
	
}