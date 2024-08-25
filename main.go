package main

import (
	"enigma/core"
	"enigma/core/utils"
	"fmt"
)

func main() {
	fmt.Println("Intializing Engima Machine")
	engima := core.InitEnigma()
	fmt.Println(engima.Cipher(utils.MsgInput()))

	// fmt.Println(engima.Cipher("ABC"))
}
