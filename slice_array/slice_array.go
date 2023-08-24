package main

import "fmt"

func main(){
	var slice1 = make([]int, 3, 3)
	slice2 := append(slice1, 4)
	slice1[0] = 100

	fmt.Println(slice1, len(slice1), cap(slice1))
	fmt.Println(slice2, len(slice2), cap(slice2))
}