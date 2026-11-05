package main

import "fmt"

func main(){
	for i := 1; i <= 10; i ++{
		for j := 1; j <= i; j++{
			fmt.Print("*")
		}//the inner loop controls the column so first it will start with i which is one
		fmt.Println()//Drops down to the next line after finishing each row
	}
}
