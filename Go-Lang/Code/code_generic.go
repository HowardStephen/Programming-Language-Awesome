package main

import "fmt"

type Element[T any] struct {
	value T
	next  *Element[T]
}

type List[T any] struct {
	head, tail *Element[T]
}

func (l *List[T]) Push(v T) {
	e := &Element[T]{v, nil}
	if l.tail == nil {
		l.head = e
		l.tail = l.head
	} else {
		l.tail.next = e
		l.tail = l.tail.next
	}
}

func (l *List[T]) Pop() (v T) {
	e := l.head
	for e.next != l.tail {
		e = e.next
	}
	v = l.tail.value
	l.tail = e
	return
}

func (l *List[T]) Traverse() {
	for e := l.head; e != l.tail.next; e = e.next {
		fmt.Print(e.value, " ")
	}
	fmt.Println()
}

func main() {
	l := List[int]{}
	l.Push(1)
	l.Push(2)
	l.Push(3)
	l.Push(4)

	l.Traverse()
	fmt.Println(l.Pop())

	l.Traverse()
}
