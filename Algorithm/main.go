package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	w := &Wait{}
	w.Add(1)
	fmt.Printf("goroutines: %d\n", runtime.NumGoroutine())
	b := w.WaitWithTimeout(5 * time.Second)
	fmt.Println(b)
	//w.Done()
	time.Sleep(10 * time.Second)
	fmt.Printf("goroutines: %d\n", runtime.NumGoroutine())
}

// Wait is similar with sync.WaitGroup which can wait with timeout
type Wait struct {
	wg sync.WaitGroup
}

// Add adds delta, which may be negative, to the WaitGroup counter.
func (w *Wait) Add(delta int) {
	w.wg.Add(delta)
}

// Done decrements the WaitGroup counter by one
func (w *Wait) Done() {
	w.wg.Done()
}

// Wait blocks until the WaitGroup counter is zero.
func (w *Wait) Wait() {
	w.wg.Wait()
}

// WaitWithTimeout blocks until the WaitGroup counter is zero or timeout
// returns true if timeout
func (w *Wait) WaitWithTimeout(timeout time.Duration) bool {
	c := make(chan bool)
	go func() {
		defer close(c)
		fmt.Println("wait...")
		w.wg.Wait()
		fmt.Println("done...")
		c <- true
	}()
	select {
	case <-c:
		return false // completed normally
	case <-time.After(timeout):
		return true // timed out
	}
}
