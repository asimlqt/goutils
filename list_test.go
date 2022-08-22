package goutils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	l := New[int]()

	exp := "goutils.List[int]"
	got := fmt.Sprintf("%T", l)

	if got != exp {
		t.Errorf("expected %s, got %s", exp, got)
	}
}

func TestAdd(t *testing.T) {
	t.Run("it returns len 1 when adding a single item", func(t *testing.T) {
		t.Parallel()

		l := New[int]()
		l.Add(2)

		exp := 1
		got := len(l)

		if got != exp {
			t.Errorf("expected list to have length of %d, got %d", exp, got)
		}
	})

	t.Run("it returns len 3 when adding 3 item items", func(t *testing.T) {
		t.Parallel()

		l := New[int]()
		l.Add(2)
		l.Add(4)
		l.Add(6)

		exp := 3
		got := len(l)

		if got != exp {
			t.Errorf("expected list to have length of %d, got %d", exp, got)
		}
	})

	t.Run("it returns len 0 when no items are added", func(t *testing.T) {
		t.Parallel()

		l := New[int]()

		exp := 0
		got := len(l)

		if got != exp {
			t.Errorf("expected list to have length of %d, got %d", exp, got)
		}
	})
}

func TestAddAll(t *testing.T) {
	t.Run("it adds all elements of a slice to the list", func(t *testing.T) {
		t.Parallel()

		l := New[int]()
		l.AddAll([]int{2, 4, 6})

		exp := 3
		got := len(l)

		if got != exp {
			t.Errorf("expected list to have length of %d, got %d", exp, got)
		}
	})

	t.Run("it adds all elements of a list to the list", func(t *testing.T) {
		t.Parallel()

		l := New[int]()

		l2 := New[int]()
		l2.Add(2)
		l2.Add(4)

		l.AddAll(l2)

		exp := 2
		got := len(l)

		if got != exp {
			t.Errorf("expected list to have length of %d, got %d", exp, got)
		}
	})
}

func TestChunk(t *testing.T) {
	t.Run("it chunks a list with equal elements in every chunk", func(t *testing.T) {
		t.Parallel()

		l := New[int]()
		l.AddAll([]int{2, 4, 6, 8})

		exp := []List[int]{
			{2, 4},
			{6, 8},
		}
		got := l.Chunk(2)

		if !reflect.DeepEqual(got, exp) {
			t.Errorf("expected list to have length of %v, got %v", exp, got)
		}
	})

	t.Run("it chunks a list with less elements in the last chunk", func(t *testing.T) {
		t.Parallel()

		l := New[int]()
		l.AddAll([]int{2, 4, 6, 8})

		exp := []List[int]{
			{2, 4, 6},
			{8},
		}
		got := l.Chunk(3)

		if !reflect.DeepEqual(got, exp) {
			t.Errorf("expected list to have length of %v, got %v", exp, got)
		}
	})
}

func TestContains(t *testing.T) {
	t.Run("it returns true when list contains element", func(t *testing.T) {
		t.Parallel()

		l := New[int]()
		l.AddAll([]int{2, 4, 6, 8})

		if l.Contains(6) != true {
			t.Errorf("expected true got false")
		}
	})

	t.Run("it returns false when list does not contain element", func(t *testing.T) {
		t.Parallel()

		l := New[int]()
		l.AddAll([]int{2, 4, 6, 8})

		if l.Contains(7) != false {
			t.Errorf("expected false got true")
		}
	})
}

func TestContainsAll(t *testing.T) {
	t.Run("it returns true when list contains all elements", func(t *testing.T) {
		t.Parallel()

		l := New[int]()
		l.AddAll([]int{2, 4, 6, 8})

		if l.ContainsAll([]int{4, 6}) != true {
			t.Errorf("expected true got false")
		}
	})

	t.Run("it returns false when list does not contain one of the elements", func(t *testing.T) {
		t.Parallel()

		l := New[int]()
		l.AddAll([]int{2, 4, 6, 8})

		if l.ContainsAll([]int{4, 5, 6}) != false {
			t.Errorf("expected false got true")
		}
	})
}

func TestClear(t *testing.T) {
	t.Run("it removes all elements from the list", func(t *testing.T) {
		t.Parallel()

		l := New[int]()
		l.AddAll([]int{2, 4, 6, 8})
		l.Clear()

		if len(l) != 0 {
			t.Errorf("expected an empty list but got a non-empty list")
		}
	})
}

func TestEmpty(t *testing.T) {
	t.Run("it returns true when list is empty", func(t *testing.T) {
		t.Parallel()

		l := New[int]()

		if l.Empty() != true {
			t.Errorf("expected true got false")
		}
	})

	t.Run("it returns false when list is not empty", func(t *testing.T) {
		t.Parallel()

		l := New[int]()
		l.AddAll([]int{2, 4, 6, 8})

		if l.Empty() != false {
			t.Errorf("expected false got true")
		}
	})
}

func TestFilter(t *testing.T) {
	t.Run("it filters elements using a callback and returns a new list", func(t *testing.T) {
		t.Parallel()

		l := New[int]()
		l.AddAll([]int{1, 2, 3, 4, 5, 6})

		evens := func(e int) bool {
			return e%2 == 0
		}

		exp := [3]int{2, 4, 6}
		var got [3]int
		copy(got[:], l.Filter(evens))

		if !reflect.DeepEqual(got, exp) {
			t.Errorf("expected %v, got %v", exp, got)
		}
	})
}

func TestFirst(t *testing.T) {
	t.Run("it returns the first element", func(t *testing.T) {
		t.Parallel()

		l := New[int]()
		l.AddAll([]int{2, 4, 6})

		if got, err := l.First(); err != nil {
			t.Errorf("expected 2, got %v", got)
		}
	})

	t.Run("it returns an error when list is empty", func(t *testing.T) {
		t.Parallel()

		l := New[int]()

		if _, err := l.First(); err == nil {
			t.Errorf("expected error, got nil")
		}
	})
}

func TestGet(t *testing.T) {
	t.Run("it returns an elemnt", func(t *testing.T) {
		t.Parallel()

		l := New[int]()
		l.AddAll([]int{2, 4, 6})

		exp := 4
		got, err := l.Get(1)

		if got != exp {
			t.Errorf("expected %d, got %d", exp, got)
		}

		if err != nil {
			t.Errorf("expected err to be nil, got %s", err)
		}
	})

	t.Run("it returns an error when index is invalid", func(t *testing.T) {
		t.Parallel()

		l := New[int]()
		l.AddAll([]int{2, 4, 6})

		_, err := l.Get(5)

		if err == nil {
			t.Errorf("expected err to not be nil")
		}
	})
}

func TestIndex(t *testing.T) {
	t.Run("it returns the index of an element", func(t *testing.T) {
		t.Parallel()

		l := New[int]()
		l.AddAll([]int{2, 4, 6})

		if got := l.Index(6); got != 2 {
			t.Errorf("expected 2, got %d", got)
		}
	})

	t.Run("it returns the index of the first element found", func(t *testing.T) {
		t.Parallel()

		l := New[int]()
		l.AddAll([]int{2, 4, 4, 6})

		if got := l.Index(4); got != 1 {
			t.Errorf("expected 2, got %d", got)
		}
	})

	t.Run("it returns -1 when element does not exist", func(t *testing.T) {
		t.Parallel()

		l := New[int]()
		l.AddAll([]int{2, 4, 6})

		if got := l.Index(7); got != -1 {
			t.Errorf("expected -1, got %d", got)
		}
	})
}

func TestInsert(t *testing.T) {
	t.Run("it inserts element at index", func(t *testing.T) {
		t.Parallel()

		l := New[int]()
		l.AddAll([]int{2, 4, 6})

		l.Insert(2, 5)

		exp := [4]int{2, 4, 5, 6}

		var got [4]int
		copy(got[:], l)

		if got != exp {
			t.Errorf("expected %v, got %v", exp, got)
		}
	})

	t.Run("it returns an error when inserting an element in an invalid index", func(t *testing.T) {
		t.Parallel()

		l := New[int]()
		l.AddAll([]int{2, 4, 6})

		if err := l.Insert(3, 5); err == nil {
			t.Errorf("expected an error")
		}
	})
}

func TestLast(t *testing.T) {
	t.Run("it returns the last element", func(t *testing.T) {
		t.Parallel()

		l := New[int]()
		l.AddAll([]int{2, 4, 6})

		if got, err := l.Last(); err != nil {
			t.Errorf("expected 6, got %v", got)
		}
	})

	t.Run("it returns an error when list is empty", func(t *testing.T) {
		t.Parallel()

		l := New[int]()

		if _, err := l.Last(); err == nil {
			t.Errorf("expected error, got nil")
		}
	})
}

func TestMap(t *testing.T) {
	t.Run("it maps over a list", func(t *testing.T) {
		t.Parallel()

		l := New[int]()
		l.AddAll([]int{1, 2, 3})

		double := func(e int) int {
			return e * 2
		}

		exp := [3]int{2, 4, 6}
		var got [3]int
		copy(got[:], l.Map(double))

		if !reflect.DeepEqual(got, exp) {
			t.Errorf("expected %v, got %v", exp, got)
		}
	})
}

func TestPopFirst(t *testing.T) {
	t.Run("it removes the first element form the list and returns it", func(t *testing.T) {
		t.Parallel()

		l := New[int]()
		l.AddAll([]int{2, 4, 6})

		got, err := l.PopFirst()

		if got != 2 {
			t.Errorf("expected 2, got %d", got)
		}

		if err != nil {
			t.Errorf("expected error to be nil, got %s", err)
		}

		if len(l) != 2 {
			t.Errorf("expected listto have length 2, got %d", len(l))
		}
	})

	t.Run("it returns an error if list is empty", func(t *testing.T) {
		t.Parallel()

		l := New[int]()

		_, err := l.PopFirst()

		if err == nil {
			t.Error("expected error to not be nil")
		}
	})
}

func TestPopLast(t *testing.T) {
	t.Run("it removes the last element form the list and returns it", func(t *testing.T) {
		t.Parallel()

		l := New[int]()
		l.AddAll([]int{2, 4, 6})

		got, err := l.PopLast()

		if got != 6 {
			t.Errorf("expected 6, got %d", got)
		}

		if err != nil {
			t.Errorf("expected error to be nil, got %s", err)
		}

		if len(l) != 2 {
			t.Errorf("expected listto have length 2, got %d", len(l))
		}
	})

	t.Run("it returns an error if list is empty", func(t *testing.T) {
		t.Parallel()

		l := New[int]()

		_, err := l.PopLast()

		if err == nil {
			t.Error("expected error to not be nil")
		}
	})
}

func TestRemove(t *testing.T) {
	t.Run("it returns true if element was removed", func(t *testing.T) {
		t.Parallel()

		l := New[int]()
		l.AddAll([]int{2, 4, 6})

		if l.Remove(4) != true {
			t.Errorf("expected true, got false")
		}

		exp := [2]int{2, 6}
		var got [2]int
		copy(got[:], l)

		if !reflect.DeepEqual(got, exp) {
			t.Errorf("expected %v, got %v", exp, got)
		}
	})

	t.Run("it returns false if element was not found", func(t *testing.T) {
		t.Parallel()

		l := New[int]()
		l.AddAll([]int{2, 4, 6})

		if l.Remove(3) != false {
			t.Errorf("expected false, got true")
		}

		exp := [3]int{2, 4, 6}
		var got [3]int
		copy(got[:], l)

		if !reflect.DeepEqual(got, exp) {
			t.Errorf("expected %v, got %v", exp, got)
		}
	})
}

func TestRemoveIndex(t *testing.T) {
	t.Run("it returns element at index", func(t *testing.T) {
		t.Parallel()

		l := New[int]()
		l.AddAll([]int{2, 4, 6, 8})

		got, err := l.RemoveIndex(2)

		if got != 6 {
			t.Errorf("expected 6, got %d", got)
		}

		if err != nil {
			t.Error("expected error to be nil")
		}
	})

	t.Run("it returns an error if element was not found", func(t *testing.T) {
		t.Parallel()

		l := New[int]()
		l.AddAll([]int{2, 4, 6, 8})

		got, err := l.RemoveIndex(4)

		if got != 0 {
			t.Errorf("expected 0, got %d", got)
		}

		if err == nil {
			t.Error("expected error")
		}
	})
}

func TestReplace(t *testing.T) {
	t.Run("it replaces one element with another", func(t *testing.T) {
		t.Parallel()

		l := New[int]()
		l.AddAll([]int{2, 4, 6})

		if err := l.Replace(6, 8); err != nil {
			t.Errorf("expected no error, got %s", err)
		}

		exp := [3]int{2, 4, 8}
		var got [3]int
		copy(got[:], l)

		if !reflect.DeepEqual(got, exp) {
			t.Errorf("expected %v, got %v", exp, got)
		}
	})

	t.Run("it returns an error if element was not found", func(t *testing.T) {
		t.Parallel()

		l := New[int]()
		l.AddAll([]int{2, 4, 6})

		if err := l.Replace(1, 8); err == nil {
			t.Error("expected error, got nil")
		}
	})
}

func TestReplaceIndex(t *testing.T) {
	t.Run("it replaces element at index with the supplied element", func(t *testing.T) {
		t.Parallel()

		l := New[int]()
		l.AddAll([]int{2, 4, 6})

		if err := l.ReplaceIndex(1, 3); err != nil {
			t.Errorf("expected error to be nil, got '%s'", err)
		}

		exp := [3]int{2, 3, 6}
		var got [3]int
		copy(got[:], l)

		if !reflect.DeepEqual(got, exp) {
			t.Errorf("expected %v, got %v", exp, got)
		}
	})

	t.Run("it returns an error if invalid index", func(t *testing.T) {
		t.Parallel()

		l := New[int]()
		l.AddAll([]int{2, 4, 6})

		if l.ReplaceIndex(3, 9) == nil {
			t.Error("expected err, got nil")
		}
	})
}

func TestReduce(t *testing.T) {
	t.Run("it reduces the elements of a list", func(t *testing.T) {
		t.Parallel()

		l := New[int]()
		l.AddAll([]int{1, 2, 3})

		sum := func(acc int, t int) int {
			return acc + t
		}

		got := Reduce(l, 0, sum)

		if got != 6 {
			t.Errorf("expected 6, got %d", got)
		}
	})
}
