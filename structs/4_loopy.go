package main

import "fmt"

//returns the one who got over 50
type Student struct{
	Name string
	Grade float64
}

func PassedStudents(students []Student) []Student{
	var passed []Student
	for _, student := range students {
		if student.Grade >= 50{
			passed = append(passed, student)
		}
	}
	return passed
}
			
func main(){
	students := []Student{
		{"Mary", 85},
		{"John", 42},
		{"Peter", 67},
		{"Jane", 35},
		{"David", 90},
	}

	result := PassedStudents(students)

	fmt.Println(result)
}
