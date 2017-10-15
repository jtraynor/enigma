# Enigma

A library that emulates the Enigma cipher machine.

Check out [the Wikipedia page](https://en.wikipedia.org/wiki/Enigma_machine) for a detailed explanation of how the machine
works, as well as interesting information about its commercial, diplomatic and military and usages throughout the 20th
century.

## Benchmarks
On average this library can encode at a rate of about 3,500,000 cps on my desktop.

	go test --cover --bench=. --benchmem
	BenchmarkDefault-8                       5000000               243 ns/op               0 B/op          0 allocs/op
	BenchmarkWithLaterRotors-8               5000000               250 ns/op               0 B/op          0 allocs/op
	BenchmarkWithPlugs-8                     5000000               298 ns/op               0 B/op          0 allocs/op
	BenchmarkWithLaterRotorsAndPlugs-8       5000000               300 ns/op               0 B/op          0 allocs/op
	PASS
	coverage: 100.0% of statements
	ok      github.com/jtraynor/enigma      6.670s
