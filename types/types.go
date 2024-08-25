package types

type Factsheet struct {
	RotorA         map[string]string `json:"RotorA"`
	RotorB         map[string]string `json:"RotorB"`
	RotorC         map[string]string `json:"RotorC"`
	ReflectorRotor map[string]string `json:"reflector"`
	Plugboard      map[string]string `json:"plugboard"`
}

type Key struct {
	RotorA string `json:"RotorA"`
	RotorB string `json:"RotorB"`
	RotorC string `json:"RotorC"`
}
