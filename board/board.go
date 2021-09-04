package board

import (
	"log"

	envdev "github.com/gvalkov/golang-evdev"
)

func InitializeBoard() {
	log.Println("Board class Initialization Called")
	inputDevices, err := envdev.ListInputDevices()
	if err != nil {
		log.Fatalln(err)
	}
	if len(inputDevices) <= 0 {
		log.Fatalln("Could not get the list of input devices; Check if the program has been run with sudo")
	}

	for _, dev := range inputDevices {
		log.Println("The Device name is: ", dev.Name)
	}
}
