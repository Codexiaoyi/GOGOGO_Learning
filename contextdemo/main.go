package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(3)
	go func() {
		for {
			fmt.Println("working....")
			wg.Done()
			time.Sleep(time.Second * 2)
		}
	}()
	wg.Wait()
}
