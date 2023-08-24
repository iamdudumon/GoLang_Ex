package main

import "fmt"

func main(){
	var slice = []int{1, 2, 3}
	fmt.Println(slice, len(slice))

	slice = append(slice, 4)
	fmt.Println(slice, len(slice))

	slice = append(slice, 5, 6, 7, 8, 9)
	fmt.Println(slice, len(slice))
}