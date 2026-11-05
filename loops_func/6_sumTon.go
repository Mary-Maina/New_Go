package main

import "fmt"

func sumToN(n int) int{
	sum := 0

	for i := 1; i <= n; i++{
		sum += i
	}
	return(sum)
}

func main(){
	result := sumToN(5)
	fmt.Println(result)
}
