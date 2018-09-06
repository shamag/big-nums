package main

import (
	lib "big-nums/lib"
	"fmt"
	"math/rand"
)

const N = 2000000000
const routines = 30

func main() {

	// Расчет суммы элементов слайса

	// a := make([]uint16, N, N)
	// for i := range a {
	// 	a[i] = uint16(rand.Intn(65535))
	// }
	// sum := lib.SumThroughRoutines(a, N, routines)
	// fmt.Printf("%s\n", sum.String())

	// sum2 := SumBig(a[:])
	// fmt.Printf("%s\n", sum2.String())

	// Расчет суммы последовательных значений, поступающих в канал
	ch := make(chan int64)
	go func() {
		for i := 0; i < N; i++ {
			ch <- rand.Int63()
		}
		close(ch)
	}()
	sum := lib.SumChan(N, routines, ch)
	fmt.Println("sum=", sum.String())
}
