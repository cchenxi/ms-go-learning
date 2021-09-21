package main

import "fmt"

type Person struct {
	Id        int
	FirstName string
	LastName  string
	Address   string
}

type Employee struct {
	Person
	ManagerId int
}

type Contractor struct {
	Person
	CompanyId int
}

func main() {
	employee := Employee{
		Person: Person{FirstName: "Xi"},
	}
	employee.LastName = "Chen"

	fmt.Println(employee.FirstName)
}
