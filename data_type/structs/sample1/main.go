package main

import "fmt"

type Employee struct {
	Id        int
	FirstName string
	LastName  string
	Address   string
}

func main() {
	employee := Employee{
		FirstName: "xi", LastName: "Chen", Address: "China",
	}
	fmt.Println(employee)

	employeeCopy := &employee
	employeeCopy.Id = 100
	employeeCopy.Address = "U.S.A"
	fmt.Println(employeeCopy)
}
