package main

import "fmt"

//The target is to remove duplicates in a slice 
func RemoveDuplicates(nums []int) []int{
	seen := make(map[int]bool)//Map is better in tracking what we have seen

	var unique []int //The final slice that will hold the unique values

	for i := 0; i < len(nums); i++{
		val := nums[i]

		if !seen[val]{
			seen[val] = true
			unique = append(unique, val)
		}
	}
	return unique
}

func main() {
    input := []int{2, 3, 2, 4, 3, 5, 2}

    output := RemoveDuplicates(input)

    fmt.Println(output) // Output: [2 3 4 5]
}
