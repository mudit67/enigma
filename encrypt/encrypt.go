package encrypt

import (
	"enigma/types"
	"fmt"
	"strings"
	// "strings"
)

var Key = "abc"

var Message = "go is awesome"

func handleRotorAdvancement(num int, inc int) int {
	num += inc
	if num > 26 {
		return num - 26
	}
	return num
}
func getRotorsConfig(initial types.RotorsConfig) func() types.RotorsConfig {
	var rotation2, rotation3 = 0, 0
	return func() types.RotorsConfig {
		initial.Rotors3 = handleRotorAdvancement(initial.Rotors3, 1)
		rotation3 = handleRotorAdvancement(rotation3, 1)
		if rotation3 == 26 {
			initial.Rotors2 = handleRotorAdvancement(initial.Rotors2, 1)
			rotation2 = handleRotorAdvancement(rotation2, 1)
			if rotation2 == 26 {
				initial.Rotors1 = handleRotorAdvancement(initial.Rotors1, 1)
			}
		}
		return initial
	}
}

func alphaFallback(alpha rune) rune {
	if alpha > 'Z' {
		alpha -= 26
	} else if alpha < 'A' {
		alpha += 26
	}
	return alpha
}

func encryptChar(key types.RotorsConfig, char rune, wiring types.RotorsWiring) (codedChar string) {
	var charIndex rune = char

	charIndex = wiring.Rotor3[alphaFallback(charIndex+rune(key.Rotors3))]
	charIndex = wiring.Rotor2[alphaFallback(charIndex+rune(key.Rotors2))]
	charIndex = wiring.Rotor1[alphaFallback(charIndex+rune(key.Rotors1))]
	charIndex = wiring.Reflector[charIndex]

	charIndex = wiring.Rotor1[alphaFallback(charIndex-rune(key.Rotors1))]

	charIndex = wiring.Rotor2[alphaFallback(charIndex-rune(key.Rotors2))]

	charIndex = wiring.Rotor3[alphaFallback(charIndex-rune(key.Rotors3))]

	codedChar = string(charIndex)

	return
}

func Encrypt(key types.RotorsConfig, message string, wiring types.RotorsWiring) string {
	message = strings.ToUpper(message)
	fmt.Print("Encrypting: ", message, " : ")

	var code string
	incrementRotors := getRotorsConfig(key)
	for _, msgChar := range message {
		code += encryptChar(key, msgChar, wiring)
		key = incrementRotors()
	}
	return code
}
