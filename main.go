package main

import (
	"enigma/core"
	"enigma/core/utils"
	"fmt"
)

func main() {
	fmt.Println("Intializing Engima Machine")
	var engima core.Enigma = utils.InitEnigma()

	println(engima.Cipher('A'))
}
