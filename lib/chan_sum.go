package lib

import (
	"math/big"
	"sync"
)

var wg sync.WaitGroup
var wg2 sync.WaitGroup

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
	// инициализипуем waitgroup - можно wg.Add(length)
	for i = 0; i < length; i++ {
		wg.Add(1)
	}
	// запуск горутин
	for i = 0; i < routinesNum; i++ {
		go SumChanSingle(ch, out)
	}
	// ждем завершения подсчета, закрываем канал
	wg.Wait()
	close(ch)
	wg2.Wait()
	// считаем итог
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
	// подсчет частичной суммы, берем значение из канала, добавляем к сохраненной сумме
	for num := range ch {
		sum.Add(&sum, big.NewInt(num))
		wg.Done()
	}
	// перед завершением отправляем результат
	defer func() {
		wg2.Done()
		out <- sum
	}()
}
