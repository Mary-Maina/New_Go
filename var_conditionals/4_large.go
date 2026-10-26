package main

import "fmt"

func large(x int, y int, z int) {
	if x > y && x > z {
		fmt.Printf("%v is the largest\n", x)
	}else if y > x && y > z {
		fmt.Printf("%v is the largest\n", y)
	}else if z > x && z > y {
		fmt.Printf("%v is the largest\n", z)
	}else if z == x && x > y {
		fmt.Printf("%v and %v are the largest\n", z, x)
	}else if x == y && y > z {
		fmt.Printf("%v and %v are the largest\n", x, y)
	}else if y == z && y > x {
		fmt.Printf("%v and %v are the largest\n", y, z)
	}else {
		fmt.Print("All numbers are equal")
	}
}

func main(){
	var x, y, z int

	fmt.Print("Please enter your numbers: ")

	_, err := fmt.Scanln(&x, &y, &z)

	if err != nil {
		fmt.Println("Invalid numbers")
		return
	}
	large(x, y, z)
}
