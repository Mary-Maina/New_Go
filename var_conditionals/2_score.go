package main
/*grading system*/

import "fmt"

func main(){
	var grade int

	fmt.Print("Input your grade: ")

	_, err := fmt.Scanln(&grade)

	if err != nil {
		fmt.Println("Invalid grade, please enter a number")
		return
	}

	if grade > 100 || grade < 0 {
		fmt.Println("Invalid grade")
	}else if grade >= 80 && grade < 100 {
		fmt.Println("Score is A")
	}else if grade >= 70 && grade < 79 {
		fmt.Println("Score is B")
	}else if grade >= 60 && grade < 69 {
		fmt.Println("Score is C")
	}else if grade >= 50 && grade < 59 {
		fmt.Println("Score is D")
	}else {
		fmt.Println("fail")
	}
}
