package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)
	go Walk(t1, c1)
	go Walk(t2, c2)

	for i := 0; i < 10; i++ {
		v1 := <-c1
		v2 := <-c2
		if v1 != v2 {
			return false
		}
	}
	return true
}

func main() {
	// test Walk()
	ch := make(chan int)
	go Walk(tree.New(7), ch)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}

	// test Same()
	t1 := tree.New(10)
	t2 := tree.New(10)
	same := Same(t1, t2)
	fmt.Println(t1)
	fmt.Println(t2)
	fmt.Println("=>", same)
}
