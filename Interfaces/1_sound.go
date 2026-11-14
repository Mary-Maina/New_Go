package main

import "fmt"

type Animal interface{
	Sound()
}

type Dog struct{}

func (d Dog) Sound(){
	fmt.Println("Dog says Woof")
}

type Cat struct{}

func (c Cat) Sound(){
	fmt.Println("Cat says Meow")
}

func MakeSound(a Animal){
	a.Sound()
}

func main(){
	cat := Cat{}
	dog := Dog{}

	MakeSound(dog)
	MakeSound(cat)
}
