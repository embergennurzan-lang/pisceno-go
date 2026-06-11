package main

import "fmt"

// Ваша функция, которая собирает алфавит
func PrintAlphabets() string {
    result := ""
    for i := 'A'; i <= 'Z'; i++ {
        result += string(i)
    }
    return result
}

// Главная функция, которая запустит код
func main() {
    // Печатаем то, что вернула функция PrintAlphabets
    fmt.Println(PrintAlphabets())
}
