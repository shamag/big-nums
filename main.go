package main

import (
	lib "big-nums/lib"
	"fmt"
	"math/rand"
)

const N = 300
const routines = 299

func main() {

	// a := make([]uint16, N, N)
	// for i := range a {
	// 	a[i] = uint16(rand.Intn(65535))
	// }
	// sum := lib.SumThroughRoutines(a, N, routines)
	// fmt.Printf("%s\n", sum.String())

	// sum2 := SumBig(a[:])
	// fmt.Printf("%s\n", sum2.String())
	ch := make(chan int64)
	go func() {
		for i := 0; i < N; i++ {
			ch <- rand.Int63()
			// ch <- 2
		}
	}()
	sum := lib.SumChan(N, routines, ch)
	fmt.Println("sum=", sum.String())
}
