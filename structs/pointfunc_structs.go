package main

import "fmt"

//we can use pointers in custom struct functions
type Student struct{
	Name string
	Age int
	Course string
	Grade float64
}

func ChangeAge(student *Student) {//*receives a pointer from student
    student.Age = 30
}

func main(){
	student1 := Student{"Mary", 22, "Microservices", 95.5}

	ChangeAge(&student1)//gives memory address of student1. 

	fmt.Println(student1.Age)
}
