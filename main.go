package main

import (
	"fmt"
	"math"
	"math/big"
	"math/rand"
)

func sumBig(sl []big.Int) big.Int {
	var sum big.Int
	sum.SetInt64(0)
	for i := range sl {
		sum.Add(&sum, &sl[i])
	}
	return sum
}
func sumBigChan(sl []big.Int, ch chan big.Int) {
	var sum big.Int
	sum.SetInt64(0)
	for i := range sl {
		sum.Add(&sum, &sl[i])
	}
	ch <- sum
}
func main() {
	const N = 200000
	const routines = 30
	var sum big.Int
	var chunkSize = float64(N) / float64((routines - 1))
	sum.SetInt64(0)

	var sumChan = make(chan big.Int)
	// sem := make(chan bool, routines)
	a := make([]big.Int, N, N)
	for i := range a {
		a[i] = *a[i].SetInt64(rand.Int63())
	}

	for chunkNum := 0; chunkNum < routines; chunkNum++ {
		end := int(math.Floor(chunkSize * float64((chunkNum + 1))))
		start := int(math.Floor(chunkSize * float64(chunkNum)))
		if end >= len(a) {
			go sumBigChan(a[start:], sumChan)
			break
		}
		go sumBigChan(a[start:end], sumChan)
	}
	var count = 1
	for part := range sumChan {
		count++
		sum.Add(&sum, &part)
		if count == routines {
			close(sumChan)
		}
	}
	sum2 := sumBig(a[:])
	fmt.Printf("%s %s \n", sum.String(), sum2.String())
}
