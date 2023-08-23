package main

import (
	"fmt"
	"bufio"
	"os"
)

func main(){
	stdin := bufio.NewReader(os.Stdin)

	var a int
	var b int
	
	n, err := fmt.Scan(&a, &b)

	if err != nil{
		fmt.Println(n, err)
		stdin.ReadString('\n')
	}else{
		fmt.Println(a,b)
	}

	fmt.Scan(&a, &b)
	fmt.Println(a,b)
}