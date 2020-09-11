package intset

import (
	"math/rand"
	"testing"
)

func TestIntSet_Add(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)
	y := []uint64{4398046511618, 0, 65536}
	for i, word := range x.words {
		if word != y[i] {
			t.Errorf("add err")
			return
		}
	}

}
func TestIntSet_Has(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)
	tests := []int{1, 144, 9, 42}
	for _, test := range tests {
		if x.Has(test) == false {
			t.Errorf("has err")
		}
	}
	if x.Has(10) == true {
		t.Errorf("has err")
	}
}
func TestIntSet_String(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)
	if x.String() != "{1 9 42 144}" {
		t.Errorf("string err:%s", x.String())
	}
}
func TestIntSet_UnionWith(t *testing.T) {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)

	y.Add(42)
	y.Add(9)
	x.UnionWith(&y)

	if x.String() != "{1 9 42 144}" {
		t.Errorf("unionwith err:%s", x.String())
	}

}
func BenchmarkIntSet_Add(b *testing.B) {
	var x IntSet
	for i := 0; i < b.N; i++ {
		x.Add(rand.Int())
	}

}
func BenchmarkIntSet_UnionWith(b *testing.B) {
	var x, y IntSet
	for i := 0; i < b.N; i++ {
		x.Add(rand.Int())
		y.Add(rand.Int())
		x.UnionWith(&y)
	}

}
