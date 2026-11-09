package main

import "fmt"

//moving all negatives to the front
func MoveNegatives(nums []int){
	write := 0

	for i := 0; i < len(nums);i++{
		if nums[i] < 0{
			nums[write], nums[i] = nums[i], nums[write]
			write++
		}
	}
}
func main() {
    nums := []int{0, -1, -5, 3, 12, -7, 6}

    MoveNegatives(nums)

    fmt.Println(nums)
}
