package utils

import (
	"encoding/json"
	"enigma/core"
	"fmt"
	"io/ioutil"
	"os"
)

func LoadJson(en *core.Enigma) {

	var jsonFile, errF = os.Open("./../../assets/config.json")

	if errF != nil {
		fmt.Errorf(errF.Error())

	}
	defer jsonFile.Close()

	byteFile, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteFile, en)
}

func InitEnigma() core.Enigma {
	var en core.Enigma
	LoadJson(&en)
	return en
}
