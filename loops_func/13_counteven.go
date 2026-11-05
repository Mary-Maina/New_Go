package main

import "fmt"

/**
countEven that accepts an integer n and returns the number of even integers between 1 and n, inclusive.
**/

func countEven(n int) int{
	sum := 0//This will count the number of times we have seen even numbers

	for i := 1; i <= n; i++{
		if i % 2 == 0{
			sum++//condition if it is even sum is added by one
		}else{
			continue //jumps the number if it is odd
		}
	}
	return sum
}
func main(){
	fmt.Println(countEven(10))
	fmt.Println(countEven(7))
	fmt.Println(countEven(1))
}
