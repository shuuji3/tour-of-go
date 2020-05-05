package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {
		fmt.Printf("type of i is %T\n", i)
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
