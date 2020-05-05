package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("(%v, %T)\n", v, v)
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("(%v, %T)\n", v, v)
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("(%v, %T)\n", v, v)
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
