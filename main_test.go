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

// func BenchmarkRoutine(b *testing.B) {
// 	for i := range a {
// 		a[i] = uint16(rand.Intn(65535))
// 	}
// 	b.Run("routines25", func(b *testing.B) {
// 		lib.SumThroughRoutines(a, N2, 25)
// 	})
// 	b.Run("routines4", func(b *testing.B) {
// 		lib.SumThroughRoutines(a, N2, 4)
// 	})
// 	b.Run("routines3", func(b *testing.B) {
// 		lib.SumThroughRoutines(a, N2, 3)
// 	})
// 	b.Run("regular", func(b *testing.B) {
// 		lib.SumBig(a[:])
// 	})
// }

func TestChanNull(t *testing.T) {
	ch := make(chan int64)
	go func() {
		for i := 0; i < 100; i++ {
			ch <- 0
		}
	}()
	sum := lib.SumChan(100, 20, ch)
	if sum.String() != big.NewInt(0).String() {
		t.Error("Expected 0, got ", sum)
	}
}

func TestChan100(t *testing.T) {
	ch := make(chan int64)
	go func() {
		for i := 0; i < 100; i++ {
			ch <- 100
		}
	}()
	sum := lib.SumChan(100, 20, ch)
	if sum.String() != big.NewInt(100*100).String() {
		t.Error("Expected 10 000, got ", sum.String())
	}
}
func sumBig2(sl []int64) big.Int {
	var sum big.Int
	sum.SetInt64(0)
	for i := range sl {
		sum.Add(&sum, big.NewInt(sl[i]))
	}
	return sum
}
func TestChanRandom(t *testing.T) {
	ch := make(chan int64)
	var count int64 = 2000000
	a := make([]int64, count, count)
	var i int64
	for i = 0; i < count; i++ {
		a[i] = rand.Int63()
	}
	go func() {
		var i int64
		for i = 0; i < count; i++ {
			ch <- a[i]
		}
	}()
	sum2 := sumBig2(a)
	sum := lib.SumChan(count, 20, ch)
	if sum.String() != sum2.String() {
		t.Error("Expected values doesnt match, got ", sum.String())
	}
}

func BenchmarkChanRoutine(b *testing.B) {
	ch := make(chan int64)
	var count int64 = 20000000
	go func() {
		var i int64
		for i = 0; i < count; i++ {
			ch <- rand.Int63()
		}
	}()
	lib.SumChan(count, 20, ch)
}
