package configuration

import (
	"errors"
	"log"
	"os"
	"path/filepath"
)

func InitializeConfig() bool {
	log.Println("Initializing Configuation")
	_folderPath := "/etc/MacroKeys/"
	// _folderPath := "/home/manjith/etc/MacroKeys/"
	_fileName := "MacroKeys.conf"
	isInitialized := _checkElseCreateConfFile(_folderPath, _fileName)

	return isInitialized
}

func _createConfigDirectory(a_strFolderPath string) bool {
	// TODO: Is this the right perms that we need for the folder?
	log.Println("Going to create the directory structure to hold the config file: ", a_strFolderPath)
	err := os.MkdirAll(a_strFolderPath, 0655)
	if err != nil {
		log.Fatalln("Could not make the required directories during first time set up: ", err)
		return false
	}
	return true
}

func _createConfigFile(a_strFileNameWithPath string) bool {

	log.Println("Full Path: ", a_strFileNameWithPath)
	_, err := os.Open(a_strFileNameWithPath)
	if errors.Is(err, os.ErrNotExist) {
		log.Println("Configuaration file does not exist; Going to Copy the template.")
		// path, err := os.Getwd()

		// utilities.CopyFile(path+"/configuration", path+"/manifests/ww-k8s-infrastructure/overlays/"+utilities.GetUsername()+"-ww")

	}
	return true
}

func _checkElseCreateConfFile(a_strFolderPath string, a_strFileName string) bool {
	log.Println("Checking if the config file exists", a_strFileName)
	_, err := os.Stat(a_strFolderPath)
	if os.IsNotExist(err) {
		log.Println("Folder does not exist. Will try to create the folders")
		if !_createConfigDirectory(a_strFolderPath) {
			return false
		}
	}
	_strFilePathAndName := filepath.Join(a_strFolderPath, a_strFileName)
	if !_createConfigFile(_strFilePathAndName) {
		return false
	}

	return true
}
