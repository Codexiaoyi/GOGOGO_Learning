package channel

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func OrderConcurrent() {
	wg.Add(9)
	dogCh := make(chan struct{}, 1)
	catCh := make(chan struct{})
	fishCh := make(chan struct{})
	dogCh <- struct{}{}
	for i := 0; i < 3; i++ {
		go Dog(dogCh, catCh)
		go Cat(catCh, fishCh)
		go Fish(fishCh, dogCh)
	}
	wg.Wait()
}

func Dog(dog <-chan struct{}, cat chan<- struct{}) {
	<-dog
	fmt.Println("dog")
	cat <- struct{}{}
	wg.Done()
}

func Cat(cat <-chan struct{}, fish chan<- struct{}) {
	<-cat
	fmt.Println("cat")
	fish <- struct{}{}
	wg.Done()
}

func Fish(fish <-chan struct{}, dog chan<- struct{}) {
	<-fish
	fmt.Println("fish")
	dog <- struct{}{}
	wg.Done()
}
