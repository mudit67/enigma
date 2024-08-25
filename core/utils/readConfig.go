package utils

import (
	"encoding/json"
	"enigma/types"
	"io"
	"os"
)

func LoadJson() (types.Factsheet, types.Key) {
	// Loading Cofig file
	var configFile, errF = os.Open("./assets/config.json")

	if errF != nil {
		panic("Error in opening config file: " + errF.Error())
	}
	defer configFile.Close()

	configBytes, _ := io.ReadAll(configFile)

	var fts types.Factsheet
	err := json.Unmarshal(configBytes, &fts)
	if err != nil {
		panic("Unable to parse Json config" + err.Error())
	}

	// Loading Key file
	var keyFile, errK = os.Open("./assets/key.json")
	if errK != nil {
		panic("Error in opening file file: " + errK.Error())
	}

	keyBytes, _ := io.ReadAll(keyFile)

	var key types.Key

	err = json.Unmarshal(keyBytes, &key)

	if err != nil {
		panic("Unable to parse Json key" + err.Error())
	}

	return fts, key
}
