package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	Walker(t, ch)
	close(ch)
}

func Walker(t *tree.Tree, ch chan int) {
	if t != nil {
		Walker(t.Left, ch)
		ch <- t.Value
		Walker(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for x := range ch1 {
		y := <-ch2
		// fmt.Printf("%v %v\n", x, y)
		if x != y {
			return false
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	// for v := range ch {
	// 	fmt.Printf("%v ", v)
	// }
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
