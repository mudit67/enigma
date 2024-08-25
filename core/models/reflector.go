package models

import (
	"enigma/core/utils"
	"errors"
)

type Reflector struct {
	WiringCfg map[rune]rune
}

func InitReflector(cfg map[string]string, plugboard bool) (Reflector, error) {
	var ref Reflector
	ref.WiringCfg = make(map[rune]rune)

	for key, value := range cfg {
		kRune, kErr := utils.StringToRune(key)
		if kErr != nil {
			return Reflector{}, errors.New("Error in parsing entry points of rotor: " + kErr.Error())
		}
		vRune, vErr := utils.StringToRune(value)
		if vErr != nil {
			return Reflector{}, errors.New("Error in parsing exit points of rotor: " + vErr.Error())
		}
		ref.WiringCfg[kRune] = vRune
	}
	if !plugboard && len(ref.WiringCfg) != 26 {
		return Reflector{}, errors.New("incorrect rotors config")
	}
	return ref, nil
}

func (pb *Reflector) Replace(in rune) (out rune) {
	out = pb.WiringCfg[in]
	if out != 0 {
		return out
	}
	return in
}
