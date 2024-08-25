package utils

import (
	"encoding/json"
	"enigma/types"
	"io"
	"os"
)

const (
	configPath = "./assets/config.json"
	keyPath    = "./assets/key.json"
)

func ParseConfigJson() (types.Factsheet, types.Key) {
	// Loading Cofig Json
	var configFile, errF = os.Open(configPath)

	if errF != nil {
		panic("Error in opening config Json: " + errF.Error())
	}
	defer configFile.Close()

	configBytes, _ := io.ReadAll(configFile)

	var fts types.Factsheet
	err := json.Unmarshal(configBytes, &fts)
	if err != nil {
		panic("Unable to parse config Json" + err.Error())
	}

	// Loading Key Json
	var keyFile, errK = os.Open(keyPath)
	if errK != nil {
		panic("Error in opening Key json: " + errK.Error())
	}

	keyBytes, _ := io.ReadAll(keyFile)

	var key types.Key

	err = json.Unmarshal(keyBytes, &key)

	if err != nil {
		panic("Unable to parse key Json" + err.Error())
	}

	return fts, key
}
