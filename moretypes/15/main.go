package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4, 5, 6, 7)
	printSlice(s)

	s1 := []int{0, 1, 2, 3}
	printSlice(s1)

	s2 := s1[:0]
	printSlice(s2)

	s3 := s2[1:2]
	printSlice(s3)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
