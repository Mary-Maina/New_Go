package main

import "fmt"
//Maps are used for items with keys and values
func main(){
	//Maps can take items of different data types
	scores := map[string]int{
		"John" : 80,
		"Daisy" : 60,
		"Mary" : 100,
		"Ian" : 75,
		"Alex" : 45,
	}
	fmt.Println(scores)//If I want to print all
	fmt.Println(scores["Mary"])//If I want to extract the value of one element

	//we can also loop through it. We initialize k as key and v as value
	for k, v := range scores {
		fmt.Println(k,"-", v)
	}
	//if we want to delete an item
	delete(scores, "Alex")
	fmt.Println(scores)

	//We can also change an item's value
	scores["Mary"] = 94
	fmt.Println(scores)
}

