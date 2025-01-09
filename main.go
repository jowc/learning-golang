package main

import "fmt"

var major = "computing"
func main() {
	var name = "John"
	lastName := "Doe"
	var age uint8 = 25
	height := 1.75
	fmt.Println(name, lastName, age, height)
	fmt.Printf("hello %v in %v", name, major)
}