package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t != nil {
		Walk(t.Left, ch)
		ch <- t.Value
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		Walk(t1, ch1)
		close(ch1)
	}()
	go func() {
		Walk(t2, ch2)
		close(ch2)
	}()

	for {
		val1, ok1 := <-ch1
		val2, ok2 := <-ch2

		if (val1 != 0) && (val2 != 0) && (val1 != val2) {
			return false
		} else if !ok1 || !ok2 {
			return true
		}
	}
}

func main() {
	fmt.Println(Same(tree.New(3), tree.New(2)))
}
