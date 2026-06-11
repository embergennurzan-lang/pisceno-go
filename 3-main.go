package main

import "fmt"

func a() {
	result := ""
	for i := 'A'; i <= 'Z'; i++ {
		result += string(i)
	}
	fmt.Println(result)
}

func main() {
	a()
}
