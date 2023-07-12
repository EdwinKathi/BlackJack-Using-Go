package main

import "fmt"

func changeValue(val *int) {
	*val = 100
}

func main() {
	num := 42
	fmt.Println("Value of num before function call:", num)

	changeValue(&num)

	fmt.Println("Value of num after function call:", num)

	var ptr *int

	if ptr == nil {
		fmt.Println("Pointer is nil")
	}

	ptr = &num

	fmt.Println("Value of num:", num)
	fmt.Println("Address of num:", &num)
	fmt.Println("Value stored in pointer variable:", *ptr)
}
