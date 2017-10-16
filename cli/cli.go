package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/jtraynor/enigma"
)

func main() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, "Enigma cipher machine emulator.\n\nUsage:\n enigma [OPTIONS] [MESSAGE]\n\nOptions:\n")
		flag.PrintDefaults()
	}

	l := flag.String("l", "III", "The rotor to be used in the left positon. Roman numerals between I - VII.")
	lr := flag.String("lr", "1", "The ring setting of the left rotor. A number between 1 - 26.")
	ls := flag.String("ls", "A", "The start positon of the left rotor. A letter between A - Z.")

	m := flag.String("m", "II", "The rotor to be used in the middle positon. Roman numerals between I - VII.")
	mr := flag.String("mr", "1", "The ring setting of the middle rotor. A number between 1 - 26.")
	ms := flag.String("ms", "A", "The start positon of the middle rotor. A letter between A - Z.")

	r := flag.String("r", "I", "The rotor to be used in the right positon. Roman numerals between I - VII.")
	rr := flag.String("rr", "1", "The ring setting of the right rotor. A number between 1 - 26.")
	rs := flag.String("rs", "A", "The start positon of the right rotor. A letter between A - Z.")

	ref := flag.String("ref", "B", "The reflector to be used. Either B or C.")

	p := flag.String("p", "", "A comma seperated list of letter pairs. e.g. \"AB,CD,EF\".")

	flag.Parse()

	message := strings.Join(flag.Args(), " ")
	if len(message) == 0 {
		flag.Usage()
		return
	}

	leftRotor := parseRotor("Left", *l)
	leftRing := parseRing("Left", *lr)
	leftStart := parseStart("Left", *ls)

	middleRotor := parseRotor("Middle", *m)
	middleRing := parseRing("Middle", *mr)
	middleStart := parseStart("Middle", *ms)

	rightRotor := parseRotor("Right", *r)
	rightRing := parseRing("Right", *rr)
	rightStart := parseStart("Right", *rs)

	reflector := parseReflector(*ref)

	plugs := parsePlugs(*p)

	e := enigma.New()

	err := e.SetRotor("left", leftRotor, leftRing, leftStart)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Left rotor \"%s\" should be a roman numeral between I - VII.\n", leftRotor)
		os.Exit(1)
	}

	err = e.SetRotor("middle", middleRotor, middleRing, middleStart)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Middle rotor \"%s\" should be a roman numeral between I - VII.\n", middleRotor)
		os.Exit(1)
	}

	err = e.SetRotor("right", rightRotor, rightRing, rightStart)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Right rotor \"%s\" should be a roman numeral between I - VII.\n", rightRotor)
		os.Exit(1)
	}

	err = e.SetReflector(reflector)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Reflector \"%s\" should be either B or C.\n", reflector)
		os.Exit(1)
	}

	err = e.AddPlugs(plugs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Plugs \"%s\" should be a comma seperated list of letter pairs.\n", *p)
		os.Exit(1)
	}

	fmt.Println(e.Encode(message))
}

func parseRotor(position, input string) string {
	if len(input) == 0 {
		switch position {
		case "Left":
			return "III"
		case "Middle":
			return "II"
		case "Right":
		default:
			return "I"
		}
	}

	return input
}

func parseRing(position, input string) int {
	if len(input) == 0 {
		return 1
	}

	ring, err := strconv.Atoi(input)
	if err != nil || ring < 1 || ring > 26 {
		fmt.Fprintf(os.Stderr, "%s rotor ring setting \"%s\" should be a number between 1 and 26.\n", position, input)
		os.Exit(1)
	}

	return ring
}

func parseStart(position, input string) rune {
	if len(input) == 0 {
		return 'A'
	}

	start := unicode.ToUpper(rune(input[0]))
	if start < 'A' || start > 'Z' {
		fmt.Fprintf(os.Stderr, "%s rotor start position \"%s\" should be a letter between A - Z.\n", position, input)
		os.Exit(1)
	}

	return start
}

func parseReflector(input string) string {
	if len(input) == 0 {
		return "B"
	}

	return input
}

func parsePlugs(input string) []string {
	if len(input) == 0 {
		return []string{}
	}

	return strings.Split(input, ",")
}
