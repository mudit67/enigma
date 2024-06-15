package types

type RotorsConfig struct {
	Rotors1, Rotors2, Rotors3 int
}

// type RotorsMap struct {
// 	Input, Ouput [26]int
// }

type RotorsWiring struct {
	Rotor1, Rotor2, Rotor3, Reflector [26]int
}
