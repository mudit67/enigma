package core

import (
	"enigma/core/models"
)

const (
	RightToLeft = +1
	LeftToRight = -1
)

type Enigma struct {
	Rotor1         models.Rotor     `json:"RotorA"`
	Rotor2         models.Rotor     `json:"RotorB"`
	Rotor3         models.Rotor     `json:"RotorC"`
	ReflectorRotor models.Reflector `json:"reflector"`
	Plugboard      models.Reflector `json:"plugboard"`
}

func (en *Enigma) incrementRotors() {
	if en.Rotor3.Advance() {
		if en.Rotor2.Advance() {
			en.Rotor1.Advance()
		}
	}
}
func (en *Enigma) Cipher(input rune) (out rune) {
	out = en.Plugboard.Replace(input)
	out = en.Rotor3.Jumple(out, RightToLeft)
	out = en.Rotor2.Jumple(out, RightToLeft)
	out = en.Rotor1.Jumple(out, RightToLeft)
	out = en.ReflectorRotor.WiringCfg[out]
	out = en.Rotor1.Jumple(out, LeftToRight)
	out = en.Rotor2.Jumple(out, LeftToRight)
	out = en.Rotor3.Jumple(out, LeftToRight)
	out = en.Plugboard.Replace(out)
	en.incrementRotors()
	return
}
