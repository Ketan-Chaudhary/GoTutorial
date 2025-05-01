package main

import (
	"fmt"
	"sync"
)

func main() {
	// through channels multiple goroutines communicate
	fmt.Println("Channels in Golang")

	// creating a channel
	myChal := make(chan int, 2)
	wg := &sync.WaitGroup{}

	// myChal <- 5
	// fmt.Println(<-myChal)
	wg.Add(2)
	// recieve only channel
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChalOpen := <-ch
		fmt.Println(isChalOpen)
		fmt.Println(val)
		fmt.Println(<-ch)
		//fmt.Println(<-ch)
		wg.Done()
	}(myChal, wg)

	// send only channel
	go func(ch chan<- int, wg *sync.WaitGroup) {
		ch <- 5
		ch <- 6
		close(ch)
		wg.Done()
	}(myChal, wg)

	wg.Wait()
}
