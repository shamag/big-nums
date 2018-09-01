package main

import (
	lib "big-nums/lib"

	"fmt"
	"math/big"
	"math/rand"
	"testing"
)

func TestNull(t *testing.T) {
	a := make([]uint16, 10, 10)
	sum := lib.SumThroughRoutines(a, 10, 2)
	if sum.String() != big.NewInt(0).String() {
		t.Error("Expected 0, got ", sum)
	}
}

func Test100(t *testing.T) {
	a := make([]uint16, 100, 100)
	for i := range a {
		a[i] = 100
	}
	sum := lib.SumThroughRoutines(a, 100, 20)
	fmt.Println(sum.String())
	if sum.String() != big.NewInt(100*100).String() {
		t.Error("Expected 10000, got ", sum.String())
	}
}

var N2 int64 = 20000000

var a = make([]uint16, N2, N2)

func BenchmarkRoutine(b *testing.B) {
	for i := range a {
		a[i] = uint16(rand.Intn(65535))
	}
	b.Run("routines25", func(b *testing.B) {
		lib.SumThroughRoutines(a, N2, 25)
	})
	b.Run("routines4", func(b *testing.B) {
		lib.SumThroughRoutines(a, N2, 4)
	})
	b.Run("routines3", func(b *testing.B) {
		lib.SumThroughRoutines(a, N2, 3)
	})
	b.Run("regular", func(b *testing.B) {
		lib.SumBig(a[:])
	})
}
