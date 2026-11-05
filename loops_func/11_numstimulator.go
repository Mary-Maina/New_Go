package main

/**
Use a loop to increase start by 1 until it reaches target. Return how many increments were required.
**/
import "fmt"

func countAttempts(start int, target int) int {
	sum := 0 
	//The first step is checking if start is equal to target
	if start == target {
		return 0
	}//Next we check if start is larger than target
	if start > target {
		for i := start; i > target; i--{
			sum++
		}
		return -sum
	}

	for i := start; i < target; i++{
		sum++
	}
	return sum
}
func main(){
	fmt.Println(countAttempts(4, 10))
	fmt.Println(countAttempts(5, 5))
	fmt.Println(countAttempts(10, 4))
}
