package goutils

import (
	"errors"
)

var (
	ErrEmpty    = errors.New("list is empty")
	ErrIndex    = errors.New("index out of range")
	ErrNotFound = errors.New("element not found")
)

type List[T comparable] []T

func New[T comparable]() List[T] {
	return List[T]{}
}

func (l *List[T]) Add(e T) {
	*l = append(*l, e)
}

func (l *List[T]) AddAll(e []T) {
	*l = append(*l, e...)
}

func (l List[T]) Chunk(size int) []List[T] {
	var l2 []List[T]

	if size <= 0 {
		return l2
	}

	if size >= len(l) {
		l2 = append(l2, l)
		return l2
	}

	for i := 0; i < len(l); i += size {
		l2 = append(l2, l[i:l.min(i+size, len(l))])
	}

	return l2
}

func (l List[T]) Contains(e T) bool {
	return l.index(e) != -1
}

func (l List[T]) ContainsAll(e []T) bool {
	for _, el := range e {
		if l.index(el) == -1 {
			return false
		}
	}
	return true
}

func (l *List[T]) Clear() {
	*l = List[T]{}
}

func (l List[T]) Empty() bool {
	return len(l) == 0
}

func (l List[T]) Filter(f func(e T) bool) List[T] {
	l2 := List[T]{}
	for _, e := range l {
		if f(e) {
			l2.Add(e)
		}
	}
	return l2
}

func (l List[T]) First() (T, error) {
	if len(l) == 0 {
		return l.zero(), ErrEmpty
	}

	return l[0], nil
}

func (l List[T]) Get(i int) (T, error) {
	err := l.validateIndex(i)
	if err != nil {
		return l.zero(), err
	}
	return l[i], nil
}

func (l List[T]) Index(e T) int {
	return l.index(e)
}

func (l *List[T]) Insert(i int, e T) error {
	err := l.validateIndex(i)
	if err != nil {
		return err
	}
	*l = append((*l)[:i], append([]T{e}, (*l)[i:]...)...)
	return nil
}

func (l List[T]) Last() (T, error) {
	if len(l) == 0 {
		return l.zero(), ErrEmpty
	}
	return l[len(l)-1], nil
}

func (l List[T]) Map(f func(e T) T) List[T] {
	l2 := List[T]{}
	for _, e := range l {
		l2.Add(f(e))
	}
	return l2
}

func (l *List[T]) PopFirst() (T, error) {
	if len(*l) == 0 {
		return l.zero(), ErrEmpty
	}
	e := (*l)[0]
	*l = (*l)[1:]
	return e, nil
}

func (l *List[T]) PopLast() (T, error) {
	if len(*l) == 0 {
		return l.zero(), ErrEmpty
	}
	e := (*l)[len(*l)-1]
	*l = (*l)[:len(*l)-1]
	return e, nil
}

func (l *List[T]) Remove(e T) bool {
	i := l.index(e)
	if i == -1 {
		return false
	}
	*l = append((*l)[:i], (*l)[i+1:]...)
	return true
}

func (l *List[T]) RemoveIndex(i int) (T, error) {
	err := l.validateIndex(i)
	if err != nil {
		return l.zero(), err
	}
	e := (*l)[i]
	*l = append((*l)[:i], (*l)[i+1:]...)
	return e, nil
}

func (l *List[T]) Replace(e1 T, e2 T) error {
	i := l.index(e1)
	if i == -1 {
		return ErrNotFound
	}
	(*l)[i] = e2
	return nil
}

func (l *List[T]) ReplaceIndex(i int, e T) error {
	err := l.validateIndex(i)
	if err != nil {
		return err
	}
	(*l)[i] = e
	return nil
}

func Reduce[T comparable, A any](l List[T], acc A, f func(acc A, e T) A) A {
	a := acc
	for _, e := range l {
		a = f(a, e)
	}
	return a
}

// ---------------------------------------------------

func (l List[T]) min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (l List[T]) zero() T {
	var e T
	return e
}

func (l List[T]) validateIndex(i int) error {
	if len(l) == 0 {
		return ErrEmpty
	}
	if i < 0 || i >= len(l) {
		return ErrIndex
	}
	return nil
}

func (l List[T]) index(e T) int {
	for i, el := range l {
		if e == el {
			return i
		}
	}
	return -1
}
