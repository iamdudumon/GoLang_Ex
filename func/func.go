package main

import "fmt"

func add(a int, b int) int{
	return a + b
}

func main(){
	c := add(5, 10)
	fmt.Printf("%d\n", c)
}
