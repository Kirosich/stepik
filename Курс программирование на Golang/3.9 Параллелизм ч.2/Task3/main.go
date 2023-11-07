// Вам необходимо написать функцию calculator следующего вида:

// func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int
// Функция получает в качестве аргументов 3 канала, и возвращает канал типа <-chan int.

// в случае, если аргумент будет получен из канала firstChan,
// в выходной (возвращенный) канал вы должны отправить квадрат аргумента.
// в случае, если аргумент будет получен из канала secondChan,
// в выходной (возвращенный) канал вы должны отправить результат умножения аргумента на 3.
// в случае, если аргумент будет получен из канала stopChan, нужно просто завершить работу функции.
// Функция calculator должна быть неблокирующей, сразу возвращая управление.
// Ваша функция получит всего одно значение в один из каналов - получили значение, обработали его, завершили работу.

// После завершения работы необходимо освободить ресурсы, закрыв выходной канал, если вы этого не сделаете, то превысите предельное время работы.

package main

import (
	"fmt"
	"time"
)

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	ch := make(chan int)
	go func(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}, ch chan int) chan int {
		defer close(ch)
		select {
		case c := <-firstChan:
			ch <- c * c
			return ch
		case c := <-secondChan:
			ch <- c * 3
			return ch
		case <-stopChan:
			return ch
		}
	}(firstChan, secondChan, stopChan, ch)
	return ch
}

func main() {
	firstChan := make(chan int)
	secondChan := make(chan int)
	stopChan := make(chan struct{})
	go func() {
		fmt.Println(<-calculator(firstChan, secondChan, stopChan))
	}()
	//firstChan <- 3
	//secondChan <- 4
	stopChan <- struct{}{}
	time.Sleep(time.Second)

}
