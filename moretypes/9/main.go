package main

import "fmt"

func main() {
	// type is slice of []int
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)
	fmt.Printf("%T\n", q)

	// type is [6]int
	q2 := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(q2)
	fmt.Printf("%T\n", q2)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}
