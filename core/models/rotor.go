package models

import (
	"enigma/core/utils"
	"errors"
)

type Rotor struct {
	CurrentPosition rune /* (A..Z) */
	wiringCfg       map[rune]rune
}

func InitRotor(cfg map[string]string, pos string) (Rotor, error) {
	var rot Rotor
	rot.wiringCfg = make(map[rune]rune)
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
		rot.wiringCfg[kRune] = vRune
	}
	if len(rot.wiringCfg) != 26 {
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
		return cur - 26
	}
	if cur < 'A' {
		return cur + 26
	}
	return cur
}

func (rot *Rotor) Pass(inputChar rune, direction int) (outChar rune) {
	var inRotorWire rune
	inRotorWire = (inputChar + (rot.CurrentPosition - 'A'))

	inRotorWire = rotorFallback(inRotorWire)
	if direction == 1 {
		outChar = rot.wiringCfg[inRotorWire]
	} else {
		for k, v := range rot.wiringCfg {
			if v == inRotorWire {
				outChar = k
				break
			}
		}
	}
	outChar -= (rot.CurrentPosition - 'A')
	return
}
