package enigma_test

import (
	"testing"

	"github.com/jtraynor/enigma"
)

type testRotorSettings struct {
	name  string
	ring  int
	start rune
}

var encodeTests = map[string]struct {
	rotor1    testRotorSettings
	rotor2    testRotorSettings
	rotor3    testRotorSettings
	reflector string
	plugs     []string
	input     string
	expected  string
}{
	"Fox Pangram": {
		rotor1:    testRotorSettings{"VII", 15, 'K'},
		rotor2:    testRotorSettings{"II", 21, 'I'},
		rotor3:    testRotorSettings{"V", 15, 'C'},
		reflector: "B",
		plugs:     []string{"AB", "CD", "EF"},
		input:     "The quick brown fox jumps over the lazy dog",
		expected:  "UGNIJ SHAKK IHFLP WKJJZ EMYWT HUFEY CFLGU",
	},
	"Sphinx Pangram": {
		rotor1:    testRotorSettings{"VIII", 19, 'H'},
		rotor2:    testRotorSettings{"II", 9, 'D'},
		rotor3:    testRotorSettings{"IV", 4, 'Q'},
		reflector: "B",
		plugs:     []string{"KL", "MN", "OP"},
		input:     "Sphinx of black quartz, judge my vow",
		expected:  "QDJSX DDWOV GZJFH FFEJX PMIVS NAFL",
	},
	"Liquor Pangram": {
		rotor1:    testRotorSettings{"II", 14, 'I'},
		rotor2:    testRotorSettings{"III", 25, 'C'},
		rotor3:    testRotorSettings{"V", 6, 'Z'},
		reflector: "C",
		plugs:     []string{"UV", "WX", "YZ"},
		input:     "Pack my box with five dozen liquor jugs",
		expected:  "TUOJI JYBFG AZEMJ AJCSV KMOPN DNSMD AT",
	},
}

func TestEncode(t *testing.T) {
	for name, tc := range encodeTests {
		t.Run(name, func(t *testing.T) {
			e := enigma.Enigma{}

			err := e.AddRotor(tc.rotor1.name, tc.rotor1.ring, tc.rotor1.start)
			if err != nil {
				panic(err)
			}

			err = e.AddRotor(tc.rotor2.name, tc.rotor2.ring, tc.rotor2.start)
			if err != nil {
				panic(err)
			}

			err = e.AddRotor(tc.rotor3.name, tc.rotor3.ring, tc.rotor3.start)
			if err != nil {
				panic(err)
			}

			err = e.SetReflector(tc.reflector)
			if err != nil {
				panic(err)
			}

			err = e.AddPlugs(tc.plugs)
			if err != nil {
				panic(err)
			}

			result := e.Encode(tc.input)
			if result != tc.expected {
				t.Errorf("Failed %s.\nExpected: %s.\nResult:   %s.", name, tc.expected, result)
			}
		})
	}
}
