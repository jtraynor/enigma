package enigma

type rotor struct {
	alphabetRing  [26]rune
	substitutions [26]rune
	triggers      []rune
}

type rotors [3]*rotor

// Encodes a letter through each rotor in turn. If inverse is false then the letter passes from right to left. If inverse is
// true then the letter passes from left to right.
func (rotors rotors) encode(letter rune, inverse bool) rune {
	result := letter

	if inverse {
		for i := 2; i >= 0; i-- {
			result = rotors[i].inverseEncode(result)
		}
	} else {
		for i := 0; i <= 2; i++ {
			result = rotors[i].encode(result)
		}
	}

	return result
}

// Encodes a letter from the right side of the rotor to the left.
func (rotor *rotor) encode(letter rune) rune {
	result := letter

	subLetter := rotor.substitutions[letter-'A']

	for i := range rotor.alphabetRing {
		if rotor.alphabetRing[i] == subLetter {
			result = rune(i + 'A')
			break
		}
	}

	return result
}

// Encodes a letter from the left side of the rotor to the right.
func (rotor *rotor) inverseEncode(letter rune) rune {
	result := letter

	ringLetter := rotor.alphabetRing[letter-'A']

	for i := range rotor.substitutions {
		if rotor.substitutions[i] == ringLetter {
			result = rune(i + 'A')
			break
		}
	}

	return result
}

// Rotates the rotor one position.
func (rotor *rotor) rotate() {
	ring := [26]rune{}
	substitutions := [26]rune{}
	for i := 0; i < 26; i++ {
		ring[i] = rotor.alphabetRing[(i+1)%26]
		substitutions[i] = rotor.substitutions[(i+1)%26]
	}
	rotor.alphabetRing = ring
	rotor.substitutions = substitutions
}

// Rotates the right rotor and handles any subsequent rotations caused by each rotors triggers.
func (rotors rotors) rotate() {
	rightRotor := rotors[0]
	middleRotor := rotors[1]
	leftRotor := rotors[2]

	rightRotor.rotate()

	if rightRotor.checkTrigger() {
		middleRotor.rotate()
	}

	if rotors.checkDoubleStep() {
		middleRotor.rotate()
		leftRotor.rotate()
	}
}

// Checks if the letter that was in the window when this rotation began is one of this rotors triggers to rotate the next
// rotor along. Assumes this rotor has already been rotated.
func (rotor rotor) checkTrigger() bool {
	for _, trigger := range rotor.triggers {
		if trigger == rotor.alphabetRing[25] {
			return true
		}
	}

	return false
}

// Checks if this rotation will trigger a double step. If the right rotor has rotated the middle rotor last rotation and the
// middle rotor will rotate the left rotor on its next rotation then a double step occurs.
func (rotors rotors) checkDoubleStep() bool {
	rightRotor := rotors[0]
	middleRotor := rotors[1]

	rightCheck := false
	for _, trigger := range rightRotor.triggers {
		if trigger == rightRotor.alphabetRing[24] {
			rightCheck = true
		}
	}

	middleCheck := false
	for _, trigger := range middleRotor.triggers {
		if trigger == middleRotor.alphabetRing[0] {
			middleCheck = true
		}
	}

	return rightCheck && middleCheck
}

// Shifts the substituation in relation to the alphabet ring by the number of positions provided.
func (rotor *rotor) setRingPosition(position int) {
	increase := position - 1

	substitutions := [26]rune{}
	for i := 0; i < 26; i++ {
		substitutions[(i+increase)%26] = increaseLetter(rotor.substitutions[i], increase)
	}
	rotor.substitutions = substitutions
}

// Rotates the rotor so that the start position is as provided.
func (rotor *rotor) setStartPosition(start rune) {
	for i := 0; i < int(start-'A'); i++ {
		rotor.rotate()
	}
}

// Returns a letter the provided number of letters more than the letter provided. Wraps back to A after Z.
func increaseLetter(letter rune, increase int) rune {
	return 'A' + ((letter - 'A' + rune(increase)) % 26)
}
