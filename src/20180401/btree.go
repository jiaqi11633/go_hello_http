package main

import (
	"fmt"
	"math/rand"
)

// A Tree is a binary tree with inters value
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

//New returns a new, random binary tree holding the value
func New(k int) *Tree {
	var t *Tree
	for _, v := range rand.Perm(10) {
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
	} else {
		t.Right = insert(t.Right, v)
	}
	return t
}

func (t *Tree) String() string {
	if t == nil {
		return "()"
	}
	s := ""
	if t.Left != nil {
		s += t.Left.String() + ""
	}
	s += fmt.Sprint(t.Value)
	if t.Right != nil {
		s += "" + t.Right.String()
	}
	return "(" + s + ")"
}

func Walk(t *Tree, ch chan int) {
	rangeTree(t, ch)
	close(ch)
}

func rangeTree(t *Tree, ch chan int) {
	if t != nil {
		rangeTree(t.Left, ch)
		ch <- t.Value
		rangeTree(t.Right, ch)
	}
}

func Same(t1, t2 *Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i := range ch1 {
		if i != <-ch2 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("二叉树遍历比较")
	fmt.Println("打印 New（1）的值")
	var ch1 = make(chan int)
	go Walk(New(1), ch1)
	for v := range ch1 {
		fmt.Println(v)
	}
	fmt.Println("打印 New（2）的值")
	var ch2 = make(chan int)
	go Walk(New(2), ch2)
	for v := range ch2 {
		fmt.Println(v)
	}
}
