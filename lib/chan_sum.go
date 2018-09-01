package lib

import (
	"math/big"
	"sync"
)

var wg sync.WaitGroup
var wg2 sync.WaitGroup

func SumChan(length int64, routinesNum int64, ch chan int64) big.Int {
	var sum big.Int
	var i int64
	out := make(chan big.Int)
	sum.SetInt64(0)
	for i = 0; i < length; i++ {
		wg.Add(1)
	}
	for i = 0; i < routinesNum; i++ {
		go SumChanSingle(ch, out)
	}
	wg.Wait()
	close(ch)
	wg2.Wait()
	for i = 0; i < routinesNum; i++ {
		tmp := <-out
		sum.Add(&sum, &tmp)
	}
	return sum
}
func SumChanSingle(ch chan int64, out chan big.Int) {
	var sum big.Int
	wg2.Add(1)
	sum.SetInt64(0)
	for num := range ch {
		sum.Add(&sum, big.NewInt(num))
		wg.Done()
	}
	defer func() {
		wg2.Done()
		out <- sum
	}()
}
