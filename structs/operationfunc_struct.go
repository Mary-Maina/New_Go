package main

import "fmt"

//functions that perform operations in structs
type Student struct{
	Name string
	Age int
	Course string
	Grade float64
}

func (s Student) Introduce(){//Introduce is a method assosiated with with the Student type. 
	fmt.Println("Hi, my name is", s.Name)
}

func main(){
	student1 := Student{"Linda", 12, "Maths", 84.9}

	student1.Introduce()
	//the receiver is this way instead of Introduce(student1) so that it can be associated as a method of Student 
}
