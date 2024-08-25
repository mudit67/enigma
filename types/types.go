package types

type Factsheet struct {
	RotorA         map[string]string `json:"rotorA"`
	RotorB         map[string]string `json:"rotorB"`
	RotorC         map[string]string `json:"rotorC"`
	ReflectorRotor map[string]string `json:"reflector"`
	Plugboard      map[string]string `json:"plugboard"`
}

type Key struct {
	RotorA string `json:"rotorA"`
	RotorB string `json:"rotorB"`
	RotorC string `json:"rotorC"`
}
