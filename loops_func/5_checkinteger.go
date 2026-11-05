package main


import "fmt"

func checkNumber(i int) {
	
	if i > 0 {
		fmt.Println("Positive")
	}else if i < 0{
		fmt.Println("Negative")
	}else {
		fmt.Println("Zero")
	}
}

func main(){
	checkNumber(8)
	checkNumber(-4)
	checkNumber(0)

}

