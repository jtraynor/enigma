# Enigma CLI
A simple command-line interface for the Enigma cipher machine library.

## Installation
	go get github.com/jtraynor/enigma
	go build -o enigma.exe github.com/jtraynor/enigma/cli

## Usage
	enigma [OPTIONS] [MESSAGE]

	Options:
		-l string
			The rotor to be used in the left positon. Roman numerals between I - VII. (default "III")
		-lr string
			The ring setting of the left rotor. A number between 1 - 26. (default "1")
		-ls string
			The start positon of the left rotor. A letter between A - Z. (default "A")
		-m string
			The rotor to be used in the middle positon. Roman numerals between I - VII. (default "II")
		-mr string
			The ring setting of the middle rotor. A number between 1 - 26. (default "1")
		-ms string
			The start positon of the middle rotor. A letter between A - Z. (default "A")
		-p string
			A comma seperated list of letter pairs. e.g. "AB,CD,EF".
		-r string
			The rotor to be used in the right positon. Roman numerals between I - VII. (default "I")
		-ref string
			The reflector to be used. Either B or C. (default "B")
		-rr string
			The ring setting of the right rotor. A number between 1 - 26. (default "1")
		-rs string
			The start positon of the right rotor. A letter between A - Z. (default "A")
