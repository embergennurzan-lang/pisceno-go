package main

import "fmt"

func FishAndChips(n int) string {
	if n < 0 {
		return "error: number is negative"
	}
	if (n%2 == 0) && (n%3 == 0) {
		return "fish and chips"
	} else if n%2 == 0 {
		return "fish"
	} else if n%3 == 0 {
		return "chips"
	} else {
		return "error: non divisible"
	}
}

func main() {
	fmt.Println("Для 6:", FishAndChips(6))
	fmt.Println("Для 4:", FishAndChips(4))
	fmt.Println("Для 9:", FishAndChips(9))
	fmt.Println("Для 5:", FishAndChips(5))
}
