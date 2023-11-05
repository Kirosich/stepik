// Напишите элемент конвейера (функцию), что запоминает предыдущее значение
// и отправляет значения на следующий этап конвейера только если оно отличается от того, что пришло ранее.
// Ваша функция должна принимать два канала - inputStream и outputStream,
// в первый вы будете получать строки, во второй вы должны отправлять
// значения без повторов. В итоге в outputStream должны остаться значения,
// которые не повторяются подряд. Не забудьте закрыть канал ;)

// Функция должна называться removeDuplicates()

// Выводить или вводить ничего не нужно!

package main

import "fmt"

func removeDuplicates(inputStream chan string, outputStream chan string) {
	var str string
	defer close(outputStream)
	for value := range inputStream {
		if value != str {
			str = value
			outputStream <- str
		}
	}
}

func main() {
	inputStream := make(chan string)
	outputStream := make(chan string)

	words := []string{"word", "kek", "lol", "lol"}

	go removeDuplicates(inputStream, outputStream)
	for i := 0; i < 3; i++ {
		inputStream <- words[i]
		fmt.Println(<-outputStream)
	}
	close(inputStream)
}
