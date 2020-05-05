package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(1000 * time.Millisecond)
	boom := time.After(5000 * time.Millisecond)
	for {
		select {
		case now := <-tick:
			fmt.Println(now, "tick.")
		case now := <-boom:
			fmt.Println(now, "BOOM!")
			return
		default:
			fmt.Print(".")
			time.Sleep(100 * time.Millisecond)
		}
	}
}
