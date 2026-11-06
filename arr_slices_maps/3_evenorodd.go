package main

import "fmt"
//Returns the number of even numbers and odd numbers
func CountEvenOdd(nums []int) (int, int){
	even := 0
	odd := 0

	for i := 0; i < len(nums); i++{

		if nums[i] % 2 == 0 {
			even++
		}else {
			odd++
		}
	}
	return even, odd
}

func main(){
	Input :=  []int{1, 2, 3, 4, 5, 6}
	result,rol := CountEvenOdd(Input)

	fmt.Println(result, rol)

	Ages := []int{2, 4, 6, 8}
	res,ro := CountEvenOdd(Ages)

	fmt.Println(res, ro)
}
