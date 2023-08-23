package main

import "fmt"

func main(){
	day := "thursday"
	switch day {
	case "monday", "tuesday":
		fmt.Println("월, 화요일은 수업 가는 날.")
		// fallthrough
		// break
	case "wednesday", "thrusday":
		fmt.Println("수, 목요일은 실습 가는 날.")
	default:
		fmt.Println("수업 안 가는 날.")
	}
}
