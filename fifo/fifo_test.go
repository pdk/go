package fifo_test

import (
	"testing"

	"github.com/pdk/go/fifo"
)

func TestFIFO(t *testing.T) {

	f := fifo.New[int]()

	f.Add(1)
	f.Add(2)
	f.Add(3)

	v1, b1 := f.Next()

	if v1 != 1 {
		t.Errorf("expected 1, got %d", v1)
	}

	if !b1 {
		t.Errorf("expected b1=true, got false")
	}

	v2, b2 := f.Next()

	if v2 != 2 {
		t.Errorf("expected 2, got %d", v2)
	}

	if !b2 {
		t.Errorf("expected b2=true, got false")
	}

	v3, b3 := f.Next()

	if v3 != 3 {
		t.Errorf("expected 3, got %d", v3)
	}

	if !b3 {
		t.Errorf("expected b3=true, got false")
	}

	v4, b4 := f.Next()

	if v4 != 0 {
		t.Errorf("expected 0, got %d", v4)
	}

	if b4 {
		t.Errorf("expected b4=false, got true")
	}

	f.Add(1)

	for v, ok := f.Next(); ok; v, ok = f.Next() {
		if !ok {
			t.Errorf("expected ok=true, but got false on %d", v)
		}

		if v < 10 {
			f.Add(v + 1)
		}
	}
}

func TestRealloc(t *testing.T) {

var f fifo.FIFO[int]

	for i:=0;i<5000;i++{
		f.Add(i)
	}

	for i:=0;i<5000;i++{
		n,ok := f.Next()
		if i!= n || !ok {
			t.Errorf("expected %d, true, but got %d, %t", i, n, ok)
		}
	}

}