package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func say(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 1000)
	}
	fmt.Println("Done with ", s)
	wg.Done()
}

func main() {
	wg.Add(1)
	go say("Hello")
	wg.Add(1)
	go say("Goodbye")
	wg.Wait()

}
