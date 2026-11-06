package main

import "fmt"

func doubleNum(scores []int){
	for i := 0; i < len(scores); i++{
		scores[i] = scores[i] * 2
	}
}

func main(){
	myScores := []int{20, 45, 80, 55}
	
	doubleNum(myScores)

	fmt.Println(myScores)
}
