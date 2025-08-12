package main

import (
	"fmt"
	"time"
)

func say(msg string) {
	for i := 0; i < 3; i++ {
		fmt.Println(msg, i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go say("Hi from goroutine!") // run concurrently
	say("Hello from main!")      // runs in main goroutine
}
