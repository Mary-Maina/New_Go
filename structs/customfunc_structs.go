package main

import "fmt"

//we are learning custom funcs in structs

type Student struct{
	Name string
	Age int
	Course string
	Grade float64
}

func PrintStudent(student Student) {//the first student is the parameter and the second one is the type
    fmt.Println("Name:", student.Name)
    fmt.Println("Age:", student.Age)
    fmt.Println("Course:", student.Course)
}
func main(){
	student1 := Student{"Mary", 27, "Go", 90.0}

	PrintStudent(student1)

}
