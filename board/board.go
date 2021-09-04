package board

import (
	"log"

	envdev "github.com/gvalkov/golang-evdev"
)

func InitializeBoard() {
	// TODO: Change the board name. Also this should come from config file
	keyboardName := "Logitech G304"
	foundBoard := false
	log.Println("Board class Initialization Called")
	inputDevices, err := envdev.ListInputDevices()
	if err != nil {
		log.Fatalln(err)
	}
	if len(inputDevices) <= 0 {
		log.Fatalln("Could not get the list of input devices; Check if the program has been run with sudo")
	}

	log.Println("Going to look for keyboard with name: ", keyboardName)
	for _, dev := range inputDevices {
		if dev.Name == keyboardName {
			log.Println("The Device name is: ", dev.Name)
			foundBoard = true
		}
	}

	if !foundBoard {
		log.Fatal("Could not find the board with name: ", keyboardName)
	}
}
