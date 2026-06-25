package generics

import "fmt"

func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	prev *element[T]
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail.next.prev = lst.tail
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) Pop() (T, bool) {
	var zero T
	if lst.tail == nil {
		return zero, false
	}
	poppedValue := lst.tail.val

	if lst.head == lst.tail {
		lst.head = nil
		lst.tail = nil
		return poppedValue, true
	}

	lst.tail = lst.tail.prev
	lst.tail.next = nil

	return poppedValue, true
}

func (lst *List[T]) AllElements() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func Generics() {
	var s = []string{"foo", "bar", "zoo"}

	var i = []int{4, 3, 2, 1}

	srch := "bar"
	isrch := 3

	fmt.Println("index of", isrch, ":", SlicesIndex(i, isrch))
	fmt.Println("index of", srch, ":", SlicesIndex(s, srch))

	_ = SlicesIndex[[]string, string](s, "zoo")

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	lst.Push(27)
	fmt.Println("list:", lst.AllElements())
}
