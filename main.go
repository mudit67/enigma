package main

import (
	"enigma/encrypt"
	"enigma/types"
	"fmt"
)

func main() {

	wiring := types.RotorsWiring{
		// Rotor1:    [26]int{16, 17, 11, 15, 1, 9, 13, 8, 4, 23, 19, 22, 21, 2, 10, 7, 25, 24, 6, 20, 14, 12, 5, 26, 18, 3},
		// Rotor2:    [26]int{16, 17, 11, 15, 1, 9, 13, 8, 4, 23, 19, 22, 21, 2, 10, 7, 25, 24, 6, 20, 14, 12, 5, 26, 18, 3},
		// Rotor3:    [26]int{16, 17, 11, 15, 1, 9, 13, 8, 4, 23, 19, 22, 21, 2, 10, 7, 25, 24, 6, 20, 14, 12, 5, 26, 18, 3},
		// Reflector: [26]int{14, 24, 15, 17, 21, 22, 25, 26, 23, 20, 16, 19, 18, 1, 3, 11, 4, 13, 12, 10, 5, 6, 9, 2, 7, 8},
		Rotor1:    map[rune]rune{'A': 'A', 'B': 'B', 'C': 'C', 'D': 'D', 'E': 'E', 'F': 'F', 'G': 'G', 'H': 'H', 'I': 'I', 'J': 'J', 'K': 'K', 'L': 'L', 'M': 'M', 'N': 'N', 'O': 'O', 'P': 'P', 'Q': 'Q', 'R': 'R', 'S': 'S', 'T': 'T', 'U': 'U', 'V': 'V', 'W': 'W', 'X': 'X', 'Y': 'Y', 'Z': 'Z'},
		Rotor2:    map[rune]rune{'A': 'A', 'B': 'B', 'C': 'C', 'D': 'D', 'E': 'E', 'F': 'F', 'G': 'G', 'H': 'H', 'I': 'I', 'J': 'J', 'K': 'K', 'L': 'L', 'M': 'M', 'N': 'N', 'O': 'O', 'P': 'P', 'Q': 'Q', 'R': 'R', 'S': 'S', 'T': 'T', 'U': 'U', 'V': 'V', 'W': 'W', 'X': 'X', 'Y': 'Y', 'Z': 'Z'},
		Rotor3:    map[rune]rune{'A': 'A', 'B': 'B', 'C': 'C', 'D': 'D', 'E': 'E', 'F': 'F', 'G': 'G', 'H': 'H', 'I': 'I', 'J': 'J', 'K': 'K', 'L': 'L', 'M': 'M', 'N': 'N', 'O': 'O', 'P': 'P', 'Q': 'Q', 'R': 'R', 'S': 'S', 'T': 'T', 'U': 'U', 'V': 'V', 'W': 'W', 'X': 'X', 'Y': 'Y', 'Z': 'Z'},
		Reflector: map[rune]rune{'A': 'N', 'B': 'O', 'C': 'P', 'D': 'Q', 'E': 'R', 'F': 'S', 'G': 'T', 'H': 'U', 'I': 'V', 'J': 'W', 'K': 'X', 'L': 'Y', 'M': 'Z', 'N': 'A', 'O': 'B', 'P': 'C', 'Q': 'D', 'R': 'E', 'S': 'F', 'T': 'G', 'U': 'H', 'V': 'I', 'W': 'J', 'X': 'K', 'Y': 'L', 'Z': 'M'},
	}

	rotorsIntial := types.RotorsConfig{
		Rotors1: 0,
		Rotors2: 0,
		Rotors3: 1,
	}

	message := "ABC"

	encryptedMsg := encrypt.Encrypt(rotorsIntial, message, wiring)

	fmt.Println(encryptedMsg)
	message = encryptedMsg

	encryptedMsg = encrypt.Encrypt(rotorsIntial, message, wiring)

	fmt.Println(encryptedMsg)

}
