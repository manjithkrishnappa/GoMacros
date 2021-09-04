package main

import (
	"fmt"

	"pgs.com/gomacros/board"
)

func initialize() {
	fmt.Println("Hello World")
	board.InitializeBoard()
}

func main() {
	initialize()
}
