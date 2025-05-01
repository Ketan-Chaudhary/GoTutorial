package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

// majorly we didn't create a variable
// we should be creating a pointer to pass the reference ot the wait group
var waitGroup sync.WaitGroup
var mut sync.Mutex

var signals = []string{"test"}

func main() {
	// go greeter("Hello")
	// greeter("World")
	websitelist := []string{
		"https://youtube.com",
		"https://go.dev",
		"https://google.com",
		"https://github.com",
	}

	for _, web := range websitelist {
		go getStatusCode(web)
		waitGroup.Add(1)
	}

	waitGroup.Wait() // always goes to the end of the main function or function being called
	// make the main function wait unit all the goroutines finish their jobs
	fmt.Println(signals)
	// in this screnario this is working fine
	// but there can be cases where multiple goroutines simulatenously try to write data into same memory block
	// so for that situation we should use Mutex
	// which lock the memory
}

// func greeter(s string) {
// 	for range 6 {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer waitGroup.Done()
	res, err := http.Get(endpoint)

	if err != nil {
		log.Fatal(err)
	}

	mut.Lock()
	signals = append(signals, endpoint)
	mut.Unlock()

	fmt.Printf("%d Status code for %s :\n", res.StatusCode, endpoint)

}
