package main

import "fmt"

func main() {
    i := 1

    for i <= 5 {
        fmt.Println(i)
	i++;//we added this to prevent an infinite loop
    }
}
