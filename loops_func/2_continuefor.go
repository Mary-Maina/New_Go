package main


import "fmt"

func main(){
	for i := 0; i < 5; i++{
		if i == 3 {
			continue//Unlike break we only jump 3 in the iteration
		}
		fmt.Println(i)
	}
	fmt.Println("Loop is finished")
}
