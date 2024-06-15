package main

import (
	"enigma/encrypt"
	"enigma/types"
	"fmt"
)

func main() {

	wiring := types.RotorsWiring{
		Rotor1:    [26]int{16, 17, 11, 15, 1, 9, 13, 8, 4, 23, 19, 22, 21, 2, 10, 7, 25, 24, 6, 20, 14, 12, 5, 26, 18, 3},
		Rotor2:    [26]int{16, 17, 11, 15, 1, 9, 13, 8, 4, 23, 19, 22, 21, 2, 10, 7, 25, 24, 6, 20, 14, 12, 5, 26, 18, 3},
		Rotor3:    [26]int{16, 17, 11, 15, 1, 9, 13, 8, 4, 23, 19, 22, 21, 2, 10, 7, 25, 24, 6, 20, 14, 12, 5, 26, 18, 3},
		Reflector: [26]int{14, 24, 15, 17, 21, 22, 25, 26, 23, 20, 16, 19, 18, 1, 3, 11, 4, 13, 12, 10, 5, 6, 9, 2, 7, 8},
	}

	rotorsIntial := types.RotorsConfig{
		Rotors1: 0,
		Rotors2: 0,
		Rotors3: 1,
	}

	message := "abc"
	// message = "SMX"
	encryptedMsg := encrypt.Encrypt(rotorsIntial, message, wiring)

	fmt.Println(encryptedMsg)

}
