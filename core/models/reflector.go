package models

type Reflector struct {
	WiringCfg map[rune]rune
}

func (pb *Reflector) Replace(in rune) (out rune) {
	out = pb.WiringCfg[in]
	if out != 0 {
		return out
	}
	return in
}
