package enigma

import (
	"fmt"
	"strings"
	"unicode"
)

// The Enigma Machine
type Enigma struct {
	rotors    rotors
	reflector rotor
	plugs     plugs
}

// New returns an instance of the Enigma machine initialised with sensible defaults.
// Left Rotor: I. Middle Rotor: II. Right Rotor: III. Reflector: B. No plugs.
func New() Enigma {
	leftRotor := availableRotors["I"]
	middleRotor := availableRotors["II"]
	rightRotor := availableRotors["III"]
	reflector := availableReflectors["B"]

	return Enigma{
		rotors: rotors{
			&leftRotor,
			&middleRotor,
			&rightRotor,
		},
		reflector: reflector,
	}
}

// Encode takes the input string, encodes each letter in turn and returns the result.
// Input characters that are not letters are ignored.
func (e Enigma) Encode(input string) string {
	result := ""

	count := 0

	for _, letter := range strings.ToUpper(input) {
		if letter < 'A' || letter > 'Z' {
			continue
		}

		// Every 5 letters add a space
		if count > 0 && count%5 == 0 {
			result += " "
		}

		count++

		e.rotors.rotate()

		// Plugs on the way in
		letter = e.plugs.replace(letter)

		// Encode right to left
		letter = e.rotors.encode(letter, false)

		letter = e.reflector.encode(letter)

		// Encode left to right
		letter = e.rotors.encode(letter, true)

		// Plugs on the way out
		letter = e.plugs.replace(letter)

		result += runeToString[letter]
	}

	return result
}

// SetRotor looks up a rotor of the provided name, sets the ring and start positions, then adds the rotor to the enigma in
// the given position. Returns an error if a rotor of that name is not available or the position does not exist.
// Valid Positions: LEFT, MIDDLE, RIGHT. Available Rotors: I, II, III, IV, V, VI, VII, VIII.
func (e *Enigma) SetRotor(position, name string, ringPosition int, startPosition rune) error {
	rotor, check := availableRotors[name]
	if !check {
		return fmt.Errorf("no such rotor: %s", name)
	}

	index, check := rotorPositions[strings.ToUpper(position)]
	if !check {
		return fmt.Errorf("no such position: %s", position)
	}

	if ringPosition < 1 || ringPosition > 26 {
		return fmt.Errorf("invalid ring position: %d", ringPosition)
	}

	start := unicode.ToUpper(startPosition)
	if start < 'A' || start > 'Z' {
		return fmt.Errorf("invalid start position: %c", start)
	}

	rotor.setRingPosition(ringPosition)

	rotor.setStartPosition(start)

	e.rotors[index] = &rotor

	return nil
}

// SetReflector looks up a reflector of the provided name and adds it to the enigma.
// Returns an error if a reflector of that name is not available. Available Reflectors: B, C.
func (e *Enigma) SetReflector(name string) error {
	reflector, check := availableReflectors[name]
	if !check {
		return fmt.Errorf("no such relector: %s", name)
	}

	e.reflector = reflector

	return nil
}

// AddPlugs takes an array of 2 character input strings and adds each pair as a plug to the enigma.
func (e *Enigma) AddPlugs(inputs []string) error {
	for _, plug := range inputs {
		err := e.AddPlug(plug)
		if err != nil {
			return fmt.Errorf("failed to add plug: %v", err)
		}
	}

	return nil
}

// AddPlug takes a 2 character input string and adds the pair as a plug to the enigma.
// Returns and error if either character of the input plug is already used by an existing plug.
func (e *Enigma) AddPlug(input string) error {
	if len(input) != 2 {
		return fmt.Errorf("invalid length: %s", input)
	}

	one := unicode.ToUpper(rune(input[0]))
	two := unicode.ToUpper(rune(input[1]))

	if one < 'A' || one > 'Z' || two < 'A' || two > 'Z' {
		return fmt.Errorf("invalid plug: %s", input)
	}

	for _, p := range e.plugs {
		if one == p[0] || one == p[1] || two == p[0] || two == p[1] {
			return fmt.Errorf("duplicate plugs: %s %s", p, input)
		}
	}

	e.plugs = append(e.plugs, plug{one, two})

	return nil
}
