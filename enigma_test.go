package enigma_test

import (
	"testing"

	"github.com/jtraynor/enigma"
)

type testRotorSettings struct {
	position string
	name     string
	ring     int
	start    rune
}

var encodeTests = map[string]struct {
	leftRotor   testRotorSettings
	middleRotor testRotorSettings
	rightRotor  testRotorSettings
	reflector   string
	plugs       []string
	input       string
	expected    string
}{
	"Fox Pangram": {
		leftRotor:   testRotorSettings{"left", "V", 15, 'C'},
		middleRotor: testRotorSettings{"middle", "II", 21, 'I'},
		rightRotor:  testRotorSettings{"right", "VII", 15, 'K'},
		reflector:   "B",
		plugs:       []string{"AB", "CD", "EF"},
		input:       "The quick brown fox jumps over the lazy dog",
		expected:    "UGNIJ SHAKK IHFLP WKJJZ EMYWT HUFEY CFLGU",
	},
	"Sphinx Pangram": {
		leftRotor:   testRotorSettings{"left", "IV", 4, 'Q'},
		middleRotor: testRotorSettings{"middle", "II", 9, 'D'},
		rightRotor:  testRotorSettings{"right", "VIII", 19, 'H'},
		reflector:   "B",
		plugs:       []string{"KL", "MN", "OP"},
		input:       "Sphinx of black quartz, judge my vow",
		expected:    "QDJSX DDWOV GZJFH FFEJX PMIVS NAFL",
	},
	"Liquor Pangram": {
		leftRotor:   testRotorSettings{"left", "V", 6, 'Z'},
		middleRotor: testRotorSettings{"middle", "III", 25, 'C'},
		rightRotor:  testRotorSettings{"right", "II", 14, 'I'},
		reflector:   "C",
		plugs:       []string{"UV", "WX", "YZ"},
		input:       "Pack my box with five dozen liquor jugs",
		expected:    "TUOJI JYBFG AZEMJ AJCSV KMOPN DNSMD AT",
	},
}

func TestEncode(t *testing.T) {
	for name, tc := range encodeTests {
		t.Run(name, func(t *testing.T) {
			e := enigma.New()

			err := e.SetRotor(tc.leftRotor.position, tc.leftRotor.name, tc.leftRotor.ring, tc.leftRotor.start)
			if err != nil {
				panic(err)
			}

			err = e.SetRotor(tc.middleRotor.position, tc.middleRotor.name, tc.middleRotor.ring, tc.middleRotor.start)
			if err != nil {
				panic(err)
			}

			err = e.SetRotor(tc.rightRotor.position, tc.rightRotor.name, tc.rightRotor.ring, tc.rightRotor.start)
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

var setRotorTests = map[string]struct {
	rotor           testRotorSettings
	isErrorExpected bool
}{
	"Invalid Rotor": {
		rotor: testRotorSettings{
			name: "X",
		},
		isErrorExpected: true,
	},
	"Invalid Position": {
		rotor: testRotorSettings{
			name:     "I",
			position: "X",
		},
		isErrorExpected: true,
	},
	"Invalid Ring": {
		rotor: testRotorSettings{
			name:     "I",
			position: "LEFT",
			ring:     0,
		},
		isErrorExpected: true,
	},
	"Invalid Start": {
		rotor: testRotorSettings{
			name:     "I",
			position: "LEFT",
			ring:     1,
			start:    ' ',
		},
		isErrorExpected: true,
	},
	"Valid": {
		rotor: testRotorSettings{
			name:     "I",
			position: "LEFT",
			ring:     1,
			start:    'A',
		},
	},
}

func TestSetRotor(t *testing.T) {
	for name, tc := range setRotorTests {
		t.Run(name, func(t *testing.T) {
			e := enigma.New()

			err := e.SetRotor(tc.rotor.position, tc.rotor.name, tc.rotor.ring, tc.rotor.start)
			if tc.isErrorExpected == (err == nil) {
				t.Errorf("Failed %s. Error: %v.", name, err)
			}
		})
	}
}

var setReflectorTests = map[string]struct {
	reflectorName   string
	isErrorExpected bool
}{
	"Invalid Rotor": {
		reflectorName:   "X",
		isErrorExpected: true,
	},
	"Valid": {
		reflectorName: "B",
	},
}

func TestSetReflector(t *testing.T) {
	for name, tc := range setReflectorTests {
		t.Run(name, func(t *testing.T) {
			e := enigma.New()

			err := e.SetReflector(tc.reflectorName)
			if tc.isErrorExpected == (err == nil) {
				t.Errorf("Failed %s. Error: %v.", name, err)
			}
		})
	}
}

var addPlugTests = map[string]struct {
	plugs           []string
	isErrorExpected bool
}{
	"Invalid Length": {
		plugs:           []string{"X"},
		isErrorExpected: true,
	},
	"Invalid Letter": {
		plugs:           []string{"++"},
		isErrorExpected: true,
	},
	"Duplicate": {
		plugs:           []string{"AB", "BC"},
		isErrorExpected: true,
	},
	"Valid": {
		plugs: []string{"AB", "CD"},
	},
}

func TestAddPlugs(t *testing.T) {
	for name, tc := range addPlugTests {
		t.Run(name, func(t *testing.T) {
			e := enigma.New()

			err := e.AddPlugs(tc.plugs)
			if tc.isErrorExpected == (err == nil) {
				t.Errorf("Failed %s. Error: %v.", name, err)
			}
		})
	}
}
