package main

/*if input is odd or eve*/

import "fmt"

func main(){
	var num int 

	fmt.Print("Enter your number: ")

	_, err := fmt.Scanln(&num)

	if err != nil {
		fmt.Println("Invalid input")
		return
	}

	if num % 2 == 0 {
		fmt.Println("The number is even")
	} else {
		fmt.Println("The number is odd")
	}
}
