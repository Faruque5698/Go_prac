package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var age string
	fmt.Println("Enter your age: ")
	//fmt.Scan(&age)
	//fmt.Println("My age is", age)
	reader := bufio.NewReader(os.Stdin)
	age, _ = reader.ReadString('\n')
	//age, _ = reader.ReadString('\n') for string
	fmt.Println(age)

	//	For string use bufio reader
	// For other data type use fmt. sacn
}
