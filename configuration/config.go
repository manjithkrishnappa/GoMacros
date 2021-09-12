package configuration

import (
	"log"
)

func InitializeConfig() {
	log.Println("Initializing Configuation")
	_checkElseCreateConfFile()
}

func _checkElseCreateConfFile() {
	log.Println("Checking if the config file exists")

}
