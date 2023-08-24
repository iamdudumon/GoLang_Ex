package main

import "fmt"

func main(){
	defer fmt.Println("1. 이 출력은 함수 종료전에 실행")
	defer fmt.Println("2. 이 출력은 함수 종료전에 실행")

	for i := 0; i < 10; i++{
		fmt.Println(i)
	}
}