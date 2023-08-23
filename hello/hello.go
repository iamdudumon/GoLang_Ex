package main

import "fmt"

func main(){
	var minimunWage int = 10
	var workingHour int = 20

	var income int = minimunWage * workingHour
	var msg string = "Good Mornig!"

	fmt.Println(minimunWage, workingHour, income, msg)
}