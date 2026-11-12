package main

import "fmt"

//we are calculating total of prices
type Product struct{
	Name string
	Price float64
	Quantity int
}

func TotalInventory(products []Product) float64{
	var total float64
	total = 0.0

	for i := 0; i < len(products); i++{
		total += products[i].Price * float64(products[i].Quantity) 
	}
	return total
}
func main(){
	products := []Product{
		{"Laptop", 50000, 2},
		{"Phone", 20000, 3},
		{"Mouse", 1000, 5},
	}
	result := TotalInventory(products)
	fmt.Println(result)
}
