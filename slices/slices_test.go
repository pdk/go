package slices_test

import (
	"fmt"
	"testing"

	"github.com/pdk/go/slices"
)

func TestMap(t *testing.T) {

	a := []int{1, 2, 3}

	c := slices.Map(a, func(i int) string {
		return fmt.Sprintf("%d", i)
	})

	if c[0] != "1" {
		t.Errorf("expected first element %#v, but got %#v", "1", c[0])
	}
}

func TestReduce(t *testing.T) {

	sum := slices.Reduce([]int{1, 2, 3}, func(a, b int) int {
		return a + b
	})

	if sum != 6 {
		t.Errorf("expected 6, got %d", sum)
	}

	sum = slices.Reduce([]int{}, func(a, b int) int {
		return a + b
	})

	if sum != 0 {
		t.Errorf("expected 0, got %d", sum)
	}

	sum = slices.Sum([]int{1, 2, 3})

	if sum != 6 {
		t.Errorf("expected 6, got %d", sum)
	}
}
