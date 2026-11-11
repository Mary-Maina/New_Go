package main

import "fmt"

//An example of a struct
type Student struct {
    Name   string
    Age    int
    Course string
    Grade  float64
}//How to declare a struct. Always start the name of the struct in a capital letter
//We write type to show the data type, which is a struct (custom type). 
func main() {
    student1 := Student{
        Name:   "Mary",
        Age:    22,
        Course: "Go",
        Grade:  85.5,
    }//How to initialize the struct. If we don't initialize everything go will autofill the other places with 0, 0.0,""
    //another way of initializing is student1 := Student{"Mary", 22, "Go", 85.5}

    fmt.Println(student1)
    fmt.Println(student1.Name)
    fmt.Println(student1.Age)
    fmt.Println(student1.Course)//if we want to extract specific values. 

    student1.Age = 23//we can also modify the values
    student1.Grade = 90.0

    fmt.Println(student1.Age)
    fmt.Println(student1.Grade)
}
