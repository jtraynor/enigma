package enigma_test

import (
	"testing"

	"github.com/jtraynor/enigma"
)

var letters = []string{
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
	"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
}

func BenchmarkDefault(b *testing.B) {
	e := enigma.New()

	for n := 0; n < b.N; n++ {
		e.Encode(letters[n%26])
	}
}

func BenchmarkWithLaterRotors(b *testing.B) {
	e := enigma.New()

	err := e.SetRotor("left", "VI", 1, 'A')
	if err != nil {
		b.Fatalf("Setup Failed: %v.", err)
	}

	err = e.SetRotor("middle", "VII", 1, 'A')
	if err != nil {
		b.Fatalf("Setup Failed: %v.", err)
	}

	err = e.SetRotor("right", "VIII", 1, 'A')
	if err != nil {
		b.Fatalf("Setup Failed: %v.", err)
	}

	for n := 0; n < b.N; n++ {
		e.Encode(letters[n%26])
	}
}

func BenchmarkWithPlugs(b *testing.B) {
	e := enigma.New()

	err := e.AddPlugs([]string{"AB", "CD", "EF", "GH", "IJ", "KL", "NM", "OP", "QR", "ST", "UV", "WX", "YZ"})
	if err != nil {
		b.Fatalf("Setup Failed: %v.", err)
	}

	for n := 0; n < b.N; n++ {
		e.Encode(letters[n%26])
	}
}

func BenchmarkWithLaterRotorsAndPlugs(b *testing.B) {
	e := enigma.New()

	err := e.SetRotor("left", "VI", 1, 'A')
	if err != nil {
		b.Fatalf("Setup Failed: %v.", err)
	}

	err = e.SetRotor("middle", "VII", 1, 'A')
	if err != nil {
		b.Fatalf("Setup Failed: %v.", err)
	}

	err = e.SetRotor("right", "VIII", 1, 'A')
	if err != nil {
		b.Fatalf("Setup Failed: %v.", err)
	}

	err = e.AddPlugs([]string{"AB", "CD", "EF", "GH", "IJ", "KL", "NM", "OP", "QR", "ST", "UV", "WX", "YZ"})
	if err != nil {
		b.Fatalf("Setup Failed: %v.", err)
	}

	for n := 0; n < b.N; n++ {
		e.Encode(letters[n%26])
	}
}
