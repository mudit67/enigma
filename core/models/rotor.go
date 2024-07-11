package models

type Rotor struct {
	CurrentPosition rune /* (A..Z) */
	wiringCfg       map[rune]rune
}

func (rot *Rotor) Advance() bool {
	rot.CurrentPosition++
	if rot.CurrentPosition > 26 {
		rot.CurrentPosition = 0
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

func (rot *Rotor) Jumple(inputChar rune, direction int) (outChar rune) {
	var inRotorWire rune
	if direction < 0 {
		inRotorWire = rotorFallback(inputChar - rot.CurrentPosition)
	} else {
		inRotorWire = rotorFallback(inputChar + rot.CurrentPosition)
	}
	return rot.wiringCfg[inRotorWire]
}
