package main

import "fmt" 

/**
We want to classify if a number is positive even/Positive odd
**/
func classifyNumber(n int)(result string){
	if n > 0 && n % 2 == 0 {
		result = "Positive Even"
	}else if n > 0 && n % 2 == 1 {
		result = "Positive Odd"
	}else if n < 0 && n % 2 == 0 {
		result = "Negative Even"
	}else if n < 0 && n % 2 != 0 {
		result = "Negative Odd"
	}else {
		result = "Zero"
	}
	return//go expects return at the bottom as it tells the function when to exit. 
}

func main(){
	fmt.Println(classifyNumber(-7))
	fmt.Println(classifyNumber(0))
	fmt.Println(classifyNumber(2))
	fmt.Println(classifyNumber(7))
}
