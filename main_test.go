package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"testing"
)

func TestNull(t *testing.T) {
	a := make([]int64, 10, 10)
	sum := SumThroughRoutines(a, 10, 2)
	if sum.String() != big.NewInt(0).String() {
		t.Error("Expected 0, got ", sum)
	}
}

func Test100(t *testing.T) {
	a := make([]int64, 100, 100)
	for i := range a {
		a[i] = 100
	}
	sum := SumThroughRoutines(a, 100, 2)
	fmt.Println(sum.String())
	if sum.String() != big.NewInt(100*100).String() {
		t.Error("Expected 10000, got ", sum.String())
	}
}

var a = make([]int64, 10000000, 10000000)

func BenchmarkRoutine(b *testing.B) {
	for i := range a {
		a[i] = rand.Int63()
	}
	SumThroughRoutines(a, 10000000, 30)
}

func BenchmarkRegular(b *testing.B) {
	for i := range a {
		a[i] = rand.Int63()
	}
	SumBig(a)
}
