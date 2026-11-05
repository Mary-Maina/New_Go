package main

import "fmt"

func printTable(n int) {
	mut := 1
	for i := 1; i <= 10; i++ {
		mut = n * i 
		fmt.Printf("%v * %v = %v\n", i, n, mut)
	}
}

func main(){
	printTable(4)
	printTable(0)
}
