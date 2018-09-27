package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func foo(c chan int, x int) {
	defer wg.Done()
	// pipe result into the channel that we passed it
	c <- x * 5
}

func main() {
	someChannel := make(chan int, 10) // make a channel that returns ints

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go foo(someChannel, i)
	}

	wg.Wait()
	close(someChannel)

	for item := range someChannel {
		fmt.Println(item)
	}

	time.Sleep(time.Second * 2)
}
