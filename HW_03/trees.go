package main

import (
	"fmt"
	"math/rand"

	//"golang.org/x/tour/tree"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func New(n, k int) *Tree {
	var t *Tree
	for _, v := range rand.Perm(n) {
		t = insert(t, (1+v)*k)
	}
	return t
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
		return t
	}
	t.Right = insert(t.Right, v)
	return t
}


// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *Tree, ch chan int) {
	Walker(t, ch)
	close(ch)
}

func Walker(t *Tree, ch chan int) {
	if t != nil {
		Walker(t.Left, ch)
		ch <- t.Value
		Walker(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *Tree) bool {
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
	go Walk(New(10, 1), ch)
	//for v := range ch {
	//	fmt.Printf("%v ", v)
	//}
	fmt.Println(Same(New(10, 1), New(10, 1)))
	fmt.Println(Same(New(10, 1), New(10, 2)))
}
