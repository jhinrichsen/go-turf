// Math high precision arithmetics

package main

import (
	"math/big"
	"testing"
)

var (
	one   = big.NewInt(1)
	two   = big.NewInt(2)
	three = big.NewInt(3)
	tmp   = big.NewInt(0)
)

func TestMemberAdd(t *testing.T) {
	tmp.Add(one, two)
	if tmp.Cmp(three) != 0 {
		t.Fatalf("Expecting %d but got %d", three, tmp)
	}
}

func BenchmarkMemberAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmp.Add(one, two)
		if tmp.Cmp(three) != 0 {
			b.Fatalf("Expecting %d but got %d", three, tmp)
		}
	}
}

func BenchmarkFuncAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmp2 := tmp.Add(one, two)
		if tmp2.Cmp(three) != 0 {
			b.Fatalf("Expecting %d but got %d", three, tmp)
		}
	}
}
