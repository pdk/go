package pgl_test

import (
	"testing"

	"github.com/pdk/go/pgl"
)

func TestMax(t *testing.T) {

	a := pgl.Max(1, 2)
	if a != 2 {
		t.Errorf("expected Max to return 2, got %d", a)
	}

	b := pgl.Max(-1, 5)
	if b != 5 {
		t.Errorf("expected Max to return 5, got %d", b)
	}

	c := pgl.Max(1.0, 2.0)
	if c != 2.0 {
		t.Errorf("expected Max to return 2, got %f", c)
	}

	d := pgl.Max(-1.0, 5.0)
	if d != 5.0 {
		t.Errorf("expected Max to return 5, got %f", d)
	}

	z := pgl.Max([]int64{}...)
	if z != 0 {
		t.Errorf("expected Max to return 0, got %d", z)
	}
}
