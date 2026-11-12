package main

import "fmt"

//increasing salary by a particular percentage
type Employee struct{
	Name string
	Age int
	Salary float64
}

func IncreaseSalary(e *Employee, percentage float64){

	e.Salary = (percentage + 100)/100 * e.Salary
}

func main(){
	employee := Employee{
		Name:   "Mary",
		Age:    25,
		Salary: 50000,
	}
	IncreaseSalary(&employee, 10)
	fmt.Printf("%.2f\n", employee.Salary)
}

/**
I can also choose to import "math" and the calculation inside the IncreaseSalary
e.Salary = math.Round(e.Salary*100) / 100 which will also work and round it to 2 decimal places
**/
