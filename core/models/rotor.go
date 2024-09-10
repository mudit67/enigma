package models

import (
	"enigma/core/utils"
	"errors"
)

type Rotor struct {
	CurrentPosition rune /* (A..Z) */
	wiringCfgPos    map[rune]rune
	wiringCfgNeg    map[rune]rune
}

func InitRotor(cfg map[string]string, pos string) (Rotor, error) {
	var rot Rotor
	rot.wiringCfgPos = make(map[rune]rune)
	rot.wiringCfgNeg = make(map[rune]rune)
	var err error
	rot.CurrentPosition, err = utils.StringToRune(pos)
	if err != nil {
		return Rotor{}, errors.New("Error in parsing Position of rotor: " + err.Error())
	}
	for key, value := range cfg {
		kRune, kErr := utils.StringToRune(key)
		if kErr != nil {
			return Rotor{}, errors.New("Error in parsing entry points of rotor: " + kErr.Error())
		}
		vRune, vErr := utils.StringToRune(value)
		if vErr != nil {
			return Rotor{}, errors.New("Error in parsing exit points of rotor: " + vErr.Error())
		}
		rot.wiringCfgPos[kRune] = vRune
		rot.wiringCfgNeg[vRune] = kRune
	}
	if len(rot.wiringCfgPos) != 26 {
		return Rotor{}, errors.New("incorrect rotors config")
	}

	return rot, nil

}
func (rot *Rotor) Advance() bool {
	rot.CurrentPosition++
	if rot.CurrentPosition > 'Z' {
		rot.CurrentPosition = 'A'
		return true
	}

	return false
}

func rotorFallback(cur rune) rune {
	if cur > 'Z' {
		// return 'A' + (cur - 'Z')
		return cur - 26
	}
	if cur < 'A' {
		// return 'Z' - ('A' - cur)
		return cur + 26
	}
	return cur
}

func (rot *Rotor) Pass(inputChar rune, direction int) (outChar rune) {
	var inRotorWire rune
	inRotorWire = (inputChar + (rot.CurrentPosition - 'A'))
	inRotorWire = rotorFallback(inRotorWire)
	if direction == 1 {
		outChar = rot.wiringCfgPos[inRotorWire]
	} else {
		outChar = rot.wiringCfgNeg[inRotorWire]
	}
	outChar -= (rot.CurrentPosition - 'A')
	outChar = rotorFallback(outChar)
	return
}
