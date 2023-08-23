package main

import "fmt"

func add(a int, b int) (sum int){		// 반환 변수 명을 지을 수 있다
	sum = a + b
	return
}

func main(){
	c := add(5, 10)
	fmt.Printf("%d\n", c)
}
