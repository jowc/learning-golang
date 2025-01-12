package main

import (
	"fmt"
	"sort"
	"strings"
)

var major = "computing"

var cars [2]string

func main() {
	var name = "John"
	lastName := "Doe"
	var age uint8 = 65
	height := 1.75
	banks := []string{"kuda", "moniepoint"}
	banks = append(banks, "mercedis")
	cars = [2]string{"hp", "honda"}
	text := "some random long text i want to slice for practice purpose."
	slicedText := text[:12] + "..."
	spliceText := strings.Split(strings.ToTitle(text), " ")
	myNumbers := []int{10, 40, 73, 23, 78, 54}
	sort.Ints(myNumbers)
	sort.Strings(spliceText)
	african := true

	// while loop
	for 20 < age {
		fmt.Printf("while looping %v \n", age)
		age--
	}

	// for loop
	for i:= 20; i <= int(age); i++ {
		fmt.Printf("for looping %v \n", i)
	}

	// for in loop
	for _, value := range banks {
		if african {
			fmt.Printf("At index: %v \n", value)
		}
	}

	fmt.Println(name, lastName, age, height)
	fmt.Printf("hello %v in %v, I love %v.\n I bank with %v \n", name, major, cars[1], banks[1])
	fmt.Println(slicedText)
	fmt.Println(spliceText)
	fmt.Println(myNumbers)
}