package main

import (
	"fmt"
	"time"
)

func main() {
	defer fmt.Println(time.Now(), "world") // evaluated at first, but printed at second
	fmt.Println(time.Now(), "hello")

	//2020-04-30 14:42:06.798614 +0900 JST m=+0.000112144 hello
	//2020-04-30 14:42:06.798612 +0900 JST m=+0.000110096 world
}
