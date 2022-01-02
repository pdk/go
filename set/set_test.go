package set_test

import (
	"testing"

	"github.com/pdk/go/set"
)

func TestSet(t *testing.T) {

	s := set.New[int]()

	s.Add(1)

	if !s.Contains(1) {
		t.Errorf("expected set to contain 1, but does not")
	}

	s.Remove(1)

	if s.Contains(1) {
		t.Errorf("expected set to not contain 1, but it does")
	}

	if s.Size() != 0 {
		t.Errorf("expected set to have size 0, but got %d", s.Size())
	}

	s.Add(2)
	s.Add(2)
	s.Add(3)

	if s.Size() != 2 {
		t.Errorf("expected set to have size 2, but got %d", s.Size())
	}
}

func TestSet2(t *testing.T) {

	var s set.Set[int]

	if s.Size() != 0 {
		t.Errorf("expected set to have size 0, but got %d", s.Size())
	}

	s.Remove(0) // should not blow up on missing map

	s.Add(1)

	if !s.Contains(1) {
		t.Errorf("expected set to contain 1, but does not")
	}
}
