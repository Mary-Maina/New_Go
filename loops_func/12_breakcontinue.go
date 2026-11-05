package main

/**
Write a program that checks numbers from 1 through 30.
Skip numbers divisible by 3.
Stop the loop completely when the number reaches 20.
**/
import "fmt"

func main(){
	for i := 1; i <= 30; i++{
		if i % 3 == 0 {
			continue
		}else if i == 20 {
			break
		}
		fmt.Printf("%v\n", i)
	}
}
