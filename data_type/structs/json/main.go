package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Id        int
	FirstName string `json:"name"`
	LastName  string
	Address   string `json:"address,omitempty"`
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
	employees := []Employee{
		Employee{
			Person: Person{FirstName: "Xi", LastName: "Chen"},
		},
		Employee{
			Person: Person{FirstName: "San", LastName: "Zhang"},
		},
	}

	data, _ := json.Marshal(employees)
	fmt.Printf("%s\n", data)

	var decoded []Employee
	err := json.Unmarshal(data, &decoded)
	if err != nil {
		return
	}
	fmt.Printf("%+v", decoded)
}
