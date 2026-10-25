package main

/*checks if number is positive or negative*/

import "fmt"

func main(){
	var num int

	fmt.Print("Enter your number: ")

	_, err := fmt.Scanln(&num)

	if err != nil {
		fmt.Println("Invalid input")
		return
	}

	if num > 0 {
		fmt.Println("Positive")
	} else if num < 0 {
		fmt.Println("Negative")
	} else {
		fmt.Println("Number is 0")
	}

}
