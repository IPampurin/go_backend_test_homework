package main

import "fmt"

// печатает "Алиса получила 99.5 баллов за сдачу экзамена по математике"
func main() {
	name := "Алиса"
	score := 99.5

	fmt.Printf("%s получила %s баллов за сдачу экзамена по математике.\n", name, score)
}