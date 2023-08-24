package main

import (
	"fmt"
	"unsafe"
)

func main(){
	a := 500
	var p *int

	p = &a

	fmt.Print("p size: %d\n", unsafe.Sizeof(p))

	fmt.Printf("p value: %p\n", p)
	fmt.Printf("p가 가르키는 memory value: %d", *p)
	
	*p = 100

	fmt.Printf("a의 값: %d\n", a)
	fmt.Printf("p가 가리키는 memory value: %d\n", *p)
}