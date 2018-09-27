package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func cleanup() {
	defer wg.Done()
	if r := recover(); r != nil {
		fmt.Println("Recovered in cleanup", r) // r will be the panic
	}
}

func say(s string) {
	// defer finishing the wg until everything else
	// in this func is done
	defer cleanup()
	for i := 0; i < 3; i++ {
		time.Sleep(time.Millisecond * 1000)
		fmt.Println(s, i)
		// pretend theres an error by panicking at the third iteration
		if i == 2 {
			panic("a 2!!") // this returns in the recover() insite cleanup
		}
	}
	fmt.Println("Done with ", s)
}

func main() {
	wg.Add(1)
	go say("Hello")
	wg.Add(1)
	go say("Goodbye")
	wg.Wait()

}
