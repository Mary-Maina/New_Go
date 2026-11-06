package main 

import "fmt"

func main(){
	colors := []string{"Purple","Green","Yellow", "Blue"} 
	//Like an array, the data type is mentioned on the right. However, you don't add the number of elements to it.

	colors = append(colors, "violet")//Unlike an array, we can add elements at the very end using append

	fmt.Println(colors)
	fmt.Println(len(colors))//shows the length of slice
	//You can also use ranges as used in arrays to know the ranges of elements[:4],[1:],[1:4]
}
