package main

import "fmt"
//finding the largest value in an array
func findMax(nums []int) int{
	max := nums[0]//max is the first integer value in the array

	for i := 0; i < len(nums); i++{
		//if max is a bigger number we skip the number
		if max >= nums[i] {
			continue
		}else{
			max = nums[i]//if max is larger than the number we change it to that value
		}
	}
	return max
}

func main(){
	scores := []int{100, 50, 70, 60}

	result := findMax(scores)

	fmt.Println(result)

	ages := []int{2, 4, 7, 9, 10, 6}

	maxAge := findMax(ages)

	fmt.Println(maxAge)

	rank := []int{-2, -4, -7, -9, -10, -6}
	maxRank := findMax(rank)
	fmt.Println(maxRank)
}

/**
An alternative for the function will be
func findMax(nums []int) int{
	max = nums[0]
	//The reason we put a _ is because the range returns both index and value. so we ignore the index
	for _, v := range nums{
		if v > max {
			max = v
		}
	}
	return max
}
