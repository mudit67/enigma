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
	RotorA         models.Rotor
	RotorB         models.Rotor
	RotorC         models.Rotor
	ReflectorRotor models.Reflector
	Plugboard      models.Reflector
}

func InitEnigma() Enigma {
	var en Enigma
	cfg, key := utils.ParseConfigJson()
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
		var out rune
		out = en.Plugboard.Replace(rune(char))
		out = en.RotorC.Pass(out, RightToLeft)
		out = en.RotorB.Pass(out, RightToLeft)
		out = en.RotorA.Pass(out, RightToLeft)
		out = en.ReflectorRotor.WiringCfg[out]
		out = en.RotorA.Pass(out, LeftToRight)
		out = en.RotorB.Pass(out, LeftToRight)
		out = en.RotorC.Pass(out, LeftToRight)
		out = en.Plugboard.Replace(out)
		en.incrementRotors()
		outStr += string(out)
	}
	return
}
