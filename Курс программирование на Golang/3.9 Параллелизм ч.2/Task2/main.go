//Внутри функции main (функцию объявлять не нужно),
// вам необходимо в отдельных горутинах вызвать функцию work()
// 10 раз и дождаться результатов выполнения вызванных функций.
//Функция work() ничего не принимает и не возвращает. Пакет "sync" уже импортирован.
package main

import (
	"fmt"
	"sync"
)

func work(i int, wg *sync.WaitGroup) {
	fmt.Println(i)
	wg.Done()
}

func main() {
	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go work(i, wg)
	}
	wg.Wait()
}

// Решение
// wg := new(sync.WaitGroup)
// for i := 0; i < 10; i++ {
//     wg.Add(1)
//     go func(wg *sync.WaitGroup){
//         defer wg.Done()
//         work()
//     }(wg)
// }
// wg.Wait()
