package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var count = 0

	for count < 10 {
		var num = rand.Intn(10) + 1
		fmt.Printf("%4v", num)
		count++
	}
	fmt.Println()
	for i := 0; i < count; i++ {
		var num = rand.Intn(10) + 1
		fmt.Printf("%4v", num)
	}
}
