package main

import (
	"log"

	"github.com/gdamore/tcell"
)

func main() {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatal(err)
	}
	defer screen.Fini()

	err = screen.Init()
	if err != nil {
		log.Fatal(err)
	}

	//game loop
	running := true
	for running {
		//draw logic

		//update  logic
	}
}
