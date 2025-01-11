package main

import "fmt"

func Loop() {
	x := 5

	for i := 0; i < x; i++ {
		fmt.Printf("looping %v", i)
	}
}