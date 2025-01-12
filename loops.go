package main

import "fmt"

func About() {
	myself := "I love everything tech."
	fmt.Print(myself)
}

func Loop() {
	x := 5

	for i := 0; i < x; i++ {
		fmt.Printf("looping %v", i)
	}
}