package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)
	r := bufio.NewReader(os.Stdin)

	for {
		carte1 := getNumber(r, "What is the value of your first card?")
		carte2 := getNumber(r, "What is the value of your second card?")

		// log.Println(whatShouldIdo(carte1, carte2))
		log.Println(states[carte1-1][carte2-1])
	}
}

func whatShouldIdo(num1, num2 int) string {
	if num1 == num2 {
		return "You should split"
	}
	sum := (num1 + num2)
	switch {
	case num1 == 10 && num2 == 1, num2 == 10 && num1 == 1:
		return "Blackjack!"
	default:
		return fmt.Sprintln(sum)
	}
}
