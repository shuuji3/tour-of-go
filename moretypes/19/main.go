package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	// If make omitted, this error happens:
	// panic: assignment to entry in nil map
	m = make(map[string]Vertex)

	// Here we can use m like Python dict() variable
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}

	fmt.Println(m)
	fmt.Println(m["Bell Labs"])
}
