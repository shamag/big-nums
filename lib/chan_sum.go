package lib

import (
	"math/big"
	"sync"
)

var wg sync.WaitGroup

// Суммирование последовательности значений поступающих в канал
// length - количество элеметов
// routinesNum - количество горутин
// ch - канал, в который поступает последовательность значений
func SumChan(length int64, routinesNum int64, ch chan int64) big.Int {
	var sum big.Int
	var i int64
	// канал, в который пишем результат
	out := make(chan big.Int)
	sum.SetInt64(0)
	// запуск горутин
	for i = 0; i < routinesNum; i++ {
		go SumChanSingle(ch, out)
	}

	wg.Wait()
	// считаем итог
	for i = 0; i < routinesNum; i++ {
		tmp := <-out
		sum.Add(&sum, &tmp)
	}
	return sum
}
func SumChanSingle(ch chan int64, out chan big.Int) {
	var sum big.Int
	wg.Add(1)
	sum.SetInt64(0)
	// подсчет частичной суммы, берем значение из канала, добавляем к сохраненной сумме
	for num := range ch {
		sum.Add(&sum, big.NewInt(num))
	}
	// перед завершением отправляем результат
	defer func() {
		wg.Done()
		out <- sum
	}()
}
