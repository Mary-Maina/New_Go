package main

import "fmt"

func main(){
	scores := [5]int{50, 70, 90, 100, 40}

	fmt.Println(scores[0])//Prints elements in index 0 which is 50
	fmt.Println(scores[2])//Prints 90

	//to change element
	scores[1] = 80//changes the element 70 to 80

	fmt.Println(scores)//How to print the full array

	scores2 := scores[:2]//Prints the range of an array here we print element 0,1 50,80
	scores3 := scores[0:3]//Prints elements from index 0,2 50, 80, 90
	scores4 := scores[2:]//Prints from 2 to the end of the array. So it will print 90, 100, 40
	fmt.Println(scores2)
	fmt.Println(scores3)
	fmt.Println(scores4)
	fmt.Println(len(scores))//len is a function that checks the length of arrays and slices
}
