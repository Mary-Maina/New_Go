package main

import "fmt"

//Given two integer slices, return their intersection.
func Intersect(nums1 []int, nums2 []int) []int{
	freq := map[int]int{}
	result := []int{}

	//Checks how many times a number is repeated in nums2
	for _, num := range nums2{
		freq[num]++
	}
	//returns numbers that have more than one frequency. Then appends them to result.
	for _, num := range nums1{
		if freq[num] > 0 {
			result = append(result, num)
			freq[num]--
		}
	}
	return result
}
func main(){
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	result := Intersect(nums1, nums2)

	fmt.Println(result)
}
