package main

import (
	"fmt"
	"net"
)

func syncTest() {
	for i := 21; i < 30; i++ {
		address := fmt.Sprintf("47.106.139.187:%d", i)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			fmt.Printf("%s closed\n", address)
			continue
		}
		conn.Close()
		fmt.Printf("%s opened\n", address)
	}
}
