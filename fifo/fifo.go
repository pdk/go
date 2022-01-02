package fifo

import (
	"github.com/pdk/go/slices"
)

type FIFO[T any] struct {
	items []T
	at int
}

func New[T any]() FIFO[T] {
	return FIFO[T]{	}
}

func (fi *FIFO[T]) Add(item T) {
	fi.items = append(fi.items, item)
}

func (fi *FIFO[T]) Next() (T,bool) {
	if len(fi.items) == 0 {
		var t T
		return t, false
	}

	item, rest := slices.HeadTail(fi.items)
	fi.at++
	if fi.at >= 1000 {
		rest = slices.Copy(rest)
		fi.at = 0
	}
	fi.items = rest

	return item, true
}
