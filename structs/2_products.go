package main

import "fmt"

//we are returning the product with the highest price
type Product struct{
	Name string
	Price float64
	Quantity int
}

func MostExpensive(products []Product) Product{
	largest := products[0]

	if len(products) == 0 {
		fmt.Println("Error: No products available")
		return Product{}
	}

	for i := 1; i < len(products); i++{
		if products[i].Price > largest.Price{
			largest = products[i]
		}
	}
	return largest
}

func main(){
	products := []Product{
		{"Laptop", 80000, 5},
		{"Phone", 45000, 10},
		{"Tablet", 30000, 7},
	}

	result := MostExpensive(products)	

	fmt.Println(result)
}
