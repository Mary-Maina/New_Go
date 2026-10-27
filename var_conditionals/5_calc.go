package main

import "fmt"

func calc(op int, x, y float64) float64 {
/*T shows the type*/
	if op == 1 {
		return x + y
	} else if op == 2 {
		return x - y
	} else if op == 3 {
		return x * y
	}else if op == 4 {
		if y == 0 {
			fmt.Println("Error: Division by O")
			return 0
		}
		return x / y
	}else {
		fmt.Println("Error: Invalid operation")
		return 0
	}
}
func main(){
	result := calc(1, 14, 10)
	fmt.Println(result)
	result = calc(3, 0, 10)
	fmt.Println(result)
	result = calc(7, 10, 2)
	fmt.Println(result)
	result = calc(4, 12, 3)
	fmt.Println(result)
	result = calc(3, 4.2, 10)
	fmt.Println(result)
	//calc(2, "name", 10)
}



