package main

import "fmt"

func onlyz() rune {
    return 'z'
}

func main() {
    // fmt.Printf с флагом %c выведет символ 'z' вместо его числового кода
    fmt.Printf("%c\n", onlyz())
}
