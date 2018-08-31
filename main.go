package main

import (
	"fmt"
	"math"
	"math/big"
	"math/rand"
)

const N = 2000
const routines = 30

func sumBig(sl []int64) big.Int {
	var sum big.Int
	sum.SetInt64(0)
	for i := range sl {
		sum.Add(&sum, big.NewInt(sl[i]))
	}
	return sum
}
func sumBigChan(sl []int64, ch chan big.Int) {
	var sum big.Int
	sum.SetInt64(0)
	for i := range sl {
		sum.Add(&sum, big.NewInt(sl[i]))
	}
	ch <- sum
}
func main() {

	var sum big.Int
	var chunkSize = float64(N) / float64((routines - 1))
	sum.SetInt64(0)

	var sumChan = make(chan big.Int)
	// sem := make(chan bool, routines)
	a := make([]int64, N, N)
	for i := range a {
		a[i] = rand.Int63()
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
