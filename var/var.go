package main

import "fmt"

var g int = 10

func main(){
	var m int = 20
	{
		s := 50
		fmt.Println(m, s, g)
	}

	// m = s + 20
}