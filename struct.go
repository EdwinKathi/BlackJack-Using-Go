package main

import "fmt"

type Person struct {
	Name    string
	Age     int
	Country string
}

func main() {
	// Create a new Person struct
	person := Person{
		Name:    "John Doe",
		Age:     30,
		Country: "USA",
	}

	// Access and modify struct fields
	fmt.Println("Person's name:", person.Name)
	fmt.Println("Person's age:", person.Age)
	fmt.Println("Person's country:", person.Country)

	person.Age = 40
	fmt.Println("Updated age:", person.Age)

	// Declare an anonymous struct
	employee := struct {
		Name     string
		Age      int
		Position string
		Salary   float64
	}{
		Name:     "Jane Smith",
		Age:      35,
		Position: "Manager",
		Salary:   5000.0,
	}

	fmt.Println("Employee's name:", employee.Name)
	fmt.Println("Employee's age:", employee.Age)
	fmt.Println("Employee's position:", employee.Position)
	fmt.Println("Employee's salary:", employee.Salary)
}
