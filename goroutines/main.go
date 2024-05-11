package main

import (
	"fmt"
	"sync"
)

func main() {
	messagesPerActivation := 1
	myChan := make(chan int, messagesPerActivation)
	wg := sync.WaitGroup{}

	wg.Add(2)

	// read only
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-myChan
		fmt.Println(isChannelOpen)
		fmt.Println(val)
		wg.Done()
	}(myChan, &wg)

	// write only
	go func(ch chan<- int, wg *sync.WaitGroup) {
		// each goroutine can only send 1 msg to the channel on each activation
		myChan <- 1000
		close(myChan)
		wg.Done()
	}(myChan, &wg)

	wg.Wait()
}
