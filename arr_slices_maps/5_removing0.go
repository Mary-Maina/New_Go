package main


import "fmt"

//Puts all zeros at the end of code
func MoveZeroes(nums []int) {
	write := 0

	for i := 0; i < len(nums); i++{
		if nums[i] != 0{
			nums[write] = nums[i]
			write++
		}
	}
	for i := write; i < len(nums); i++{
		nums[i] = 0
	}
}
func main() {
    nums := []int{0, 1, 0, 3, 12}

    MoveZeroes(nums)

    fmt.Println(nums)
}
