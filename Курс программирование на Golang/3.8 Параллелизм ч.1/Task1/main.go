// Напишите функцию которая принимает канал и число N типа int. Необходимо вернуть значение N+1 в канал.
// Функция должна называться task().

// Внимание! Пакет и функция main уже объявлены, выводить и считывать ничего не нужно!

package main

import "fmt"

func task(ch chan int, N int) {
	ch <- N + 1
}

func main() {
	ch := make(chan int)
	go task(ch, 3)
	fmt.Println(<-ch)
}
