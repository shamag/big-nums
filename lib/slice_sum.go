package lib

import (
	"math"
	"math/big"
)

func SumBig(sl []uint16) big.Int {
	var sum big.Int
	sum.SetInt64(0)
	for i := range sl {
		sum.Add(&sum, big.NewInt(int64(sl[i])))
	}
	return sum
}
func SumBigChan(sl []uint16, ch chan big.Int) {
	var sum big.Int
	sum.SetInt64(0)
	for i := range sl {
		sum.Add(&sum, big.NewInt(int64(sl[i])))
	}
	ch <- sum
}

func SumThroughRoutines(a []uint16, sliceLen int64, routines int) big.Int {
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
