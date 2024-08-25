package core

import (
	"enigma/core/models"
	"enigma/core/utils"
)

const (
	RightToLeft = +1
	LeftToRight = -1
)

type Enigma struct {
	RotorA         models.Rotor     `json:"RotorA"`
	RotorB         models.Rotor     `json:"RotorB"`
	RotorC         models.Rotor     `json:"RotorC"`
	ReflectorRotor models.Reflector `json:"reflector"`
	Plugboard      models.Reflector `json:"plugboard"`
}

func InitEnigma() Enigma {
	var en Enigma
	cfg, key := utils.LoadJson()
	var err error = nil

	en.RotorA, err = models.InitRotor(cfg.RotorA, key.RotorA)
	if err != nil {
		panic("[ERROR]: Rotor A: " + err.Error())
	}

	en.RotorB, err = models.InitRotor(cfg.RotorB, key.RotorB)
	if err != nil {
		panic("[ERROR]: Rotor B: " + err.Error())
	}

	en.RotorC, err = models.InitRotor(cfg.RotorC, key.RotorC)
	if err != nil {
		panic("[ERROR]: Rotor C: " + err.Error())
	}

	en.ReflectorRotor, err = models.InitReflector(cfg.ReflectorRotor, false)
	if err != nil {
		panic("[ERROR]: Reflector: " + err.Error())
	}

	en.Plugboard, err = models.InitReflector(cfg.ReflectorRotor, true)
	if err != nil {
		panic("[ERROR]: Plugboard: " + err.Error())
	}

	return en
}
func (en *Enigma) incrementRotors() {
	if en.RotorC.Advance() {
		if en.RotorB.Advance() {
			en.RotorA.Advance()
		}
	}
}
func (en *Enigma) Cipher(input string) (outStr string) {
	for _, char := range input {
		// fmt.Println(input, char, string(char))
		var out rune
		out = en.Plugboard.Replace(rune(char))
		// fmt.Print(string(out), " ")
		out = en.RotorC.Jumble(out, RightToLeft)
		// fmt.Print(string(out), " ")
		out = en.RotorB.Jumble(out, RightToLeft)
		// fmt.Print(string(out), " ")
		out = en.RotorA.Jumble(out, RightToLeft)
		// fmt.Print(string(out), " ")
		out = en.ReflectorRotor.WiringCfg[out]
		// fmt.Print(string(out), " ")
		out = en.RotorA.Jumble(out, LeftToRight)
		// fmt.Print(string(out), " ")
		out = en.RotorB.Jumble(out, LeftToRight)
		// fmt.Print(string(out), " ")
		out = en.RotorC.Jumble(out, LeftToRight)
		// fmt.Print(string(out), " ")
		out = en.Plugboard.Replace(out)
		// fmt.Println()
		en.incrementRotors()
		outStr += string(out)
		// fmt.Println(outStr, out)
	}
	return
}
