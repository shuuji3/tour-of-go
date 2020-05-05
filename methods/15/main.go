package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64)     // panic
	f, ok = i.(float64) // ok: (f, ok) = (0, false) same as above
	fmt.Println(f)
}
