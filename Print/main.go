package main

import "fmt"

func main() {
	fmt.Println("Hello")
	fmt.Printf("Nahid \n")

	name := "Ashad"
	age := 25
	height := 5.858
	fmt.Println("My name is", name, "and My age is", age, "and My height is", height)
	fmt.Printf("My name is %s\n", name)
	fmt.Printf("and My age is %d\n", age)
	fmt.Printf("and My height is %.1f\n", height)

	fmt.Printf("Type of name is %T\n", name)
	fmt.Printf("and Type of age is %T\n", age)
	fmt.Printf("and Type of height is %T\n", height)

}
