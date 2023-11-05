// Напишите функцию которая принимает канал и строку. Необходимо отправить эту же строку
// в заданный канал 5 раз, добавив к ней пробел.

// Функция должна называться task2().

package main

import "fmt"

func task2(ch chan string, str string) {
	str = str + " "
	for i := 0; i < 5; i++ {
		ch <- str
	}
}

func main() {
	ch := make(chan string)
	go task2(ch, "kek")
	for i := 0; i < 5; i++ {
		fmt.Print(<-ch)
	}
}
