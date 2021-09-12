package main

import (
	"fmt"

	"pgs.com/gomacros/board"
	"pgs.com/gomacros/configuration"
)

func initialize() {
	fmt.Println("Hello World")
	board.InitializeBoard()
	configuration.InitializeConfig()
}

func main() {
	initialize()
}
