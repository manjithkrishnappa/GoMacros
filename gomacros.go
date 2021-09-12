package main

import (
	"fmt"
	"log"

	"pgs.com/gomacros/board"
	"pgs.com/gomacros/configuration"
)

func initialize() {
	fmt.Println("Hello World")
	if !configuration.InitializeConfig() {
		log.Fatalln("Could not initalize configuration")
	}
	board.InitializeBoard()
}

func main() {
	initialize()
}
