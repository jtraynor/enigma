package enigma

type plug [2]rune

type plugs []plug

// Checks if the input is on either side of this plug then replace it with the opposite side.
// If neither side of the plug match then the input is returned unchanged.
func (plug plug) replace(input rune) rune {
	if input == plug[0] {
		return plug[1]
	} else if input == plug[1] {
		return plug[0]
	}

	return input
}

// Calls Replace on all of the plugs in this collection and returns once one is matched.
// If no plugs match then the input is returned unchanged.
func (plugs plugs) replace(input rune) rune {
	for _, plug := range plugs {
		letter := plug.replace(input)
		if letter != input {
			return letter
		}
	}

	return input
}

func (plug plug) String() string {
	return string(plug[0]) + string(plug[1])
}
