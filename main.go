package main

import "fmt"

func main() {
	name := "John"
	age := 30

	formattedString := fmt.Sprintf("Name: %s, Age: %d", name, age)
	fmt.Println(formattedString)
}
