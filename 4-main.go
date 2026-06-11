package main

import "fmt"

// Ваша функция подсчета длины
func StrLen(s string) int {
	runeStr := []rune(s)
	return len(runeStr)
}

// Добавляем вот этот блок:
func main() {
	// Пишем любой текст для проверки
	слово := "Терминал"

	// Выводим результат на экран
	fmt.Println(StrLen(слово))
}
