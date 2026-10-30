package main

import "fmt"

func main(){
	for i := 0; i < 5; i++ {
		if i == 3 {
			break //Entirely terminates the loop and we immediately go outside it
		}
		fmt.Println(i)//continues printing until we reach 3
	}
	fmt.Println("Loop finished!")//Here we print if we are outside the loop. 
}//so we will print 0, 1, 2, loop finished
