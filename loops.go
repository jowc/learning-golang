package main

import (
	"fmt"
	"strings"
)

func about() {
	myself := "I love everything tech."
	fmt.Print(myself)
}

func loop() {
	x := 5

	for x < 5 {
		fmt.Println("x is still less than 5")
		x++
	}

	for x := 10; x > 5; x-- {
		fmt.Println("x is now", x)
	}

	for i := 0; i < x; i++ {
		fmt.Printf("looping %v", i)
	}
}

func guessName(input string){
	// convert user input to lowercase
	name := strings.ToLower(input)

	switch name {
		case "jo":
			fmt.Printf("You are close, missing a letter.")
		case "joe":
			fmt.Println("You are right.")
		default:
			fmt.Printf("Wrong guess.")
	}
}