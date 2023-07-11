package main

import "fmt"

type std struct {
	id   int
	name string
	avg  int
}

func main() {

	//Arithmetic Operators
	a := 5
	b := 10
	fmt.Println("Arithmetic Operators:")
	fmt.Println("Addition:", a+b)
	fmt.Println("Substraction:", a-b)
	fmt.Println("Multiplication:", a*b)
	fmt.Println("Division:", b/a)
	fmt.Println("Remainder:", b%a)
	fmt.Println()

	//Logical Operators
	var x bool = true
	var y bool = false
	fmt.Println("Logical Operators:")
	if !(x && y) {
		fmt.Println("AND")
	}
	if x || y {
		fmt.Println("OR")
	}
	fmt.Println()
	//Relational Operators
	c := 5
	fmt.Println("Relational Operators:")
	if a == c {
		fmt.Println("equals to")
	}
	if a != b {
		fmt.Println("Not equal to")
	}
	if a < b {
		fmt.Println("Less than")
	}
	if a > b {
		fmt.Println("greater than")
	}
	fmt.Println()

	//for loop with slice
	var days = []string{"sun", "mon", "tues", "wed", "thrus", "fri", "sat"}
	fmt.Println("for loop with slice")
	for i, j := range days {
		fmt.Println("the indes is", i, "the value is", j)
	}
	fmt.Println()

	//structure
	fmt.Println("Structure of a student")
	var s1 std
	s1.id = 1
	s1.name = "Avery"
	s1.avg = 58
	fmt.Println(s1)

	var s2 = std{2, "Ben", 67}
	fmt.Println(s2)

}