package main

import (
	lib "big-nums/lib"
	"fmt"
	"math/rand"
)

const N = 200
const routines = 30

func main() {

	a := make([]uint16, N, N)
	for i := range a {
		a[i] = uint16(rand.Intn(65535))
	}
	sum := lib.SumThroughRoutines(a, N, routines)
	fmt.Printf("%s\n", sum.String())

	// sum2 := SumBig(a[:])
	// fmt.Printf("%s\n", sum2.String())
}
