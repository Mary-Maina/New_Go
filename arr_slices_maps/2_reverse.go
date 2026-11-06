package main

import "fmt"

func Reverse(nums []int) []int{
	var rev []int
	//we are starting from the end of the loop. subtrating -1 tells go that we are starting with the last element
	for i := len(nums) - 1; i >= 0; i--{
		rev = append(rev, nums[i])
	}
	return rev
}

func main(){
	Input := []int{1, 2, 3, 4, 5}

	output := Reverse(Input)

	fmt.Println(output)
}
