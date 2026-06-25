package iterator

import (
	"fmt"
	"iter"
	"slices"
	"strings"
)

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {

		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}
}

func genFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 0, 1

		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func Iterator() {
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	lst.Push(26)
	lst.Push(28)
	lst.Push(30)

	i := 0
	for e := range lst.All() {

		fmt.Println(e)
		i++
		if i > 3 {
			break
		}
	}

	all := slices.Collect(lst.All())
	fmt.Println("all:", all)

	for part := range strings.SplitSeq("go-by-example", "-") {
		fmt.Printf("part: %s\n", part)
	}

	for n := range genFib() {

		if n >= 100 {
			break
		}
		fmt.Println(n)
	}
}
