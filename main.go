package main

import (
	"fmt"
	"math"
	"math/big"
	"math/rand"
)

const N = 59
const routines = 30

func SumBig(sl []int64) big.Int {
	var sum big.Int
	sum.SetInt64(0)
	for i := range sl {
		sum.Add(&sum, big.NewInt(sl[i]))
	}
	return sum
}
func SumBigChan(sl []int64, ch chan big.Int) {
	var sum big.Int
	sum.SetInt64(0)
	for i := range sl {
		sum.Add(&sum, big.NewInt(sl[i]))
	}
	ch <- sum
}

func SumThroughRoutines(a []int64, sliceLen int64, routines int) big.Int {
	var sum big.Int
	sum.SetInt64(0)
	var chunkSize = float64(sliceLen) / float64((routines - 1))

	var sumChan = make(chan big.Int)

	for chunkNum := 0; chunkNum < routines; chunkNum++ {
		end := int(math.Floor(chunkSize * float64((chunkNum + 1))))
		start := int(math.Floor(chunkSize * float64(chunkNum)))
		if end > len(a) {
			end = len(a)
		}
		go SumBigChan(a[start:end], sumChan)
	}
	var count = 0
	for part := range sumChan {
		count++
		sum.Add(&sum, &part)
		if count == routines {
			close(sumChan)
		}
	}
	return sum
}
func main() {

	a := make([]int64, N, N)
	for i := range a {
		a[i] = rand.Int63()
	}
	sum := SumThroughRoutines(a, N, routines)
	fmt.Printf("%s\n", sum.String())

	sum2 := SumBig(a[:])
	fmt.Printf("%s\n", sum2.String())
}
