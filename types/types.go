package types

type RotorsConfig struct {
	Rotors1, Rotors2, Rotors3 int
}
type RotorsWiring struct {
	Rotor1, Rotor2, Rotor3, Reflector map[rune]rune
}
