package main
/*we want to learn expressions and output they provide*/

import "fmt"

func main() {
	var score int = 10 //we can also use score := 10
    score += 5
    score = score * 2
    score -= 4

    fmt.Println(score)
}
