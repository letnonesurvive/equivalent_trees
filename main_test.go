package main

import (
	"fmt"
	"math"
	"testing"

	"golang.org/x/tour/tree"
)

var numbers1 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func TestWalk(t *testing.T) {
	t1 := tree.New(1)
	ch1 := make(chan int)

	go func() {
		Walk(t1, ch1)
		close(ch1)
	}()

	for i := 0; true; i++ {
		val1, ok := <-ch1
		if !ok {
			return
		}
		if val1 != numbers1[i] {
			t.Error(val1)
			t.Error(numbers1[i])
			t.Errorf("test for Walk Failed - error")
		}
	}
}

const k = 2

var numbers2 = [...]int{1 * k, 2 * k, 3 * k, 4 * k, 5 * k, 6 * k, 7 * k, 8 * k, 9 * k, 10 * k}

func TestWalkWithK(t *testing.T) {
	t1 := tree.New(k)
	ch1 := make(chan int)

	go func() {
		Walk(t1, ch1)
		close(ch1)
	}()

	for i := 0; true; i++ {
		val1, ok := <-ch1
		if !ok {
			return
		}
		if val1 != numbers2[i] {
			t.Error(val1)
			t.Error(numbers2[i])
			t.Errorf("test for WalkWithK Failed - error")
		}
	}
}

func TestSame(t *testing.T) {
	if !Same(tree.New(k), tree.New(k)) {
		t.Errorf("test for Same Failed - error")
	}
}

func TestUnsame(t *testing.T) {
	if Same(tree.New(k), tree.New(k-1)) {
		t.Errorf("test for Unsame Failed - error")
	}
}

func FuzzSameK(f *testing.F) {
	testcases := []int{2, 3, 4}
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, k int) {
		fmt.Println(k)
		if math.Abs(float64(k)) > 20 {
			return
		}
		if !Same(tree.New(k), tree.New(k)) {
			t.Errorf("test for k Failed - error")
		} else if Same(tree.New(k), tree.New(k+1)) {
			t.Errorf("test for k and k-1 Failed - error")
		}
	})
}
