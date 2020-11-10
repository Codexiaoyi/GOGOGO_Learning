package main

import (
	"fmt"
)

func main() {
	var hour = 28 * 24
	fmt.Printf("啊实打实的%4v", 56000000/hour)

	const command, exit = "我的GO", "aaa"

	fmt.Println("存在G吗？", exit, command)
}
