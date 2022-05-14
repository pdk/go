package convert_test

import (
	"testing"

	"github.com/pdk/go/convert"
)

func TestToAny(t *testing.T) {

	x := convert.ToAny(1, 2)

	if len(x) != 2 {
		t.Errorf("expected slice of len 2, got %d", len(x))
	}

	one, ok := x[0].(int)
	if !ok {
		t.Errorf("expected 1, but type assertion failed")
	}

	if one != 1 {
		t.Errorf("expected one == 1, but got %d", one)
	}
}
