package main

import "fmt"

//counts the frequency of int in a slice/array
func Frequency(nums []int) map[int]int{
	freq := map[int]int{}//How we declare an empty map

	//now we want the values from nums
	for _, num := range nums{
		freq[num]++
	}

	return freq
}
func main() {
    nums := []int{1, 2, 2, 3, 1, 1}

    result := Frequency(nums)

    fmt.Println(result)
}
