package main

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
)

func main() {
	// log.SetFlags(0)
	// r := bufio.NewReader(os.Stdin)
	g, err := gocui.NewGui(gocui.Output256)
	if err != nil {
		panic(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}

	// for {
	// 	g.
	// 	carte1 := getNumber(r, "Quelle est la valeur de votre première carte?")
	// 	carte2 := getNumber(r, "Quelle est la valeur de votre deuxième carte?")
	//
	// 	// log.Println(whatShouldIdo(carte1, carte2))
	// 	log.Println(states[carte1-1][carte2-1])
	// }
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
