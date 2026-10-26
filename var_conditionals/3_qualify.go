package main

import "fmt"

func main(){
	var age int
	var isMember bool 

	fmt.Println("Please enter your age and if you are a member: ")

	_, err := fmt.Scanln(&age, &isMember)

	if err != nil {
		fmt.Println("Invalid output")
		return
	}

	if age >= 18 && isMember == true {
		fmt.Println("Eligible")
	}else if age >= 18 && isMember == false {
		fmt.Println("Not Eligible, membership")
	}else if age < 18 && isMember == true {
		fmt.Println("Not Eligible, age")
	}else {
		fmt.Println("Not Eligible, age and not member")
	}
}


