# Enigma

A library that emulates the Enigma cipher machine.

Check out the [Wikipedia page](https://en.wikipedia.org/wiki/Enigma_machine) for a detailed explanation of how the machine
works, as well as interesting information about its commercial, diplomatic and military and usages throughout the 20th
century.

## Usage
You can use this library in 3 ways:
1. Interact with the [command-line interface](cli) directly.
2. Import this package into your own Go code.
3. Build the CLI and integrate the binary from just about any other language/environment.

There's examples of how to use this package in the [CLI code](cli/cli.go) as well as the [unit tests](enigma_test.go).

## Benchmarks
On average this library can encode at a rate of about 3,500,000
[cps](https://en.wikipedia.org/wiki/Printer_(computing)#Printing_speed) on my desktop.

	go test --cover --bench=. --benchmem
	BenchmarkDefault-8                       5000000               243 ns/op               0 B/op          0 allocs/op
	BenchmarkWithLaterRotors-8               5000000               250 ns/op               0 B/op          0 allocs/op
	BenchmarkWithPlugs-8                     5000000               298 ns/op               0 B/op          0 allocs/op
	BenchmarkWithLaterRotorsAndPlugs-8       5000000               300 ns/op               0 B/op          0 allocs/op
	PASS
	coverage: 100.0% of statements
	ok      github.com/jtraynor/enigma      6.670s
