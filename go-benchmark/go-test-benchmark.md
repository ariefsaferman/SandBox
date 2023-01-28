## Benchmark using Go Testing Packages

The go testing package contains a benchmarking feature that can be used to examine the performance of your code.
Writing a benchmark is very similar to writing a test, but there are some key differences:
* Benchmark functions start with Benchmark not Test.
* Benchmark functions are run several times by the testing package. The value of b.N will increase each time until the benchmark runner is satisfied with the stability of the benchmark. 
* Each benchmark must execute the code under test b.N times using a for loop.

Benchmark tests typically go in _test.go files and are named beginning with Benchmark. The testing runner executes each benchmark function several times, increasing b.N on each run until it collects a precise measurement.

---

This is an example project for benchmarking in go using go testing packages
In this project, we only tests 3 built in functions from golang:
- fmt.Sprintf()
- strings.builder()
- bytes.buffer()
  
These function will be benchmarked for concatenating strings to know which function **runs more efficiently**.

After benchmarking these functions you will see that eventhough these functions have the same output, they will have significant diferences in speed.

---

The second example will be implementing fibonacci function using 2 different algorithm:
- Recursive algorithm
- Sequential algorithm

These two function will give the same output but because it is using different algorithm, it will have different execution speed.

---

## Important arguments
```
go test -bench=BenchmarkConcat ./... -benchtime=10000x -count=5 -benchmem | tee current.txt
```
* `-benchtime=N` it accepts time or amount, for example `20s` (20 seconds) or `20x` (20 times)
* `-count=N` indicates how many times the benchmarks should be run
* `-benchmem` includes memory allocations statistics in the results
* to put the result of benchmark into a file, use `tee` command followed by the file name
* `-bench` used to specify the spesific benchmark function that we want to run
* `./...` will specify the directory of the test files

---

### Example Output:
```
goos: darwin
goarch: arm64
pkg: git.garena.com/sea-labs-id/tariner/nata.nael/go-benchmark/benchmarks/fibo
BenchmarkTestRecursiveFibonacci_10-10           13270371               181.8 ns/op
BenchmarkTestRecursiveFibonacci_20-10             104409             22296 ns/op
BenchmarkSequentialFibonacci_10-10              642636164                3.728 ns/op
BenchmarkSequentialFibonacci_20-10              350164257                6.834 ns/op
BenchmarkSequentialFibonacci_40-10              183331062               13.07 ns/op
BenchmarkSequentialFibonacci_1000-10             7491408               320.0 ns/op
PASS
ok      git.garena.com/sea-labs-id/tariner/nata.nael/go-benchmark/benchmarks/fibo       18.708s
```
The output format is:
`Benchmark<test-name>-<number-of-cpus> executions speed`

- **allocs/op** means how many distinct memory allocations occurred per op (single iteration).
- **B/op** is how many bytes were allocated per op

The first line of the result, **PASS**, comes from the testing portion of the test driver, asking go test to run your benchmarks does not disable the tests in the package.

Each line will show the final value of **b.N** iteration and the average run time of the function tested for the final value. 
In this case, my system can execute **BenchmarkTestRecursiveFibonacci_10** in 181.8 nano seconds, the final value of b.N iterations is 13.270.371 times, the -10 after the benchmark function test name is the amount of CPU used to run the benchmark.

In times where performance is important, being able to benchmark how your program performs and analyze where potential bottlenecks are, is really valuable. By understanding where these bottlenecks lie, we can more effectively determine where to focus our efforts in order to improve the performance of our systems.

---

## Benchstat

Through development, we want to know if the current speed has been improved or not, so we need to compare current and new reports.

Benchstat computes and compares statistics about benchmarks

```
go install golang.org/x/perf/cmd/benchstat@latest
```

```
benchstat [options] old.txt [new.txt] [more.txt ...]
```

```
benchstat current.txt new.txt 
```

To be able to use downloaded go packages from your terminal, add this to your .bashrc or .zshrc file:
```
export GOPATH="$HOME/go"
PATH="$GOPATH/bin:$PATH"
```

---

### Example Output
```
name                          old time/op  new time/op  delta
TestRecursiveFibonacci_10-10   180ns ± 0%   180ns ± 0%   ~     (p=1.000 n=1+1)
TestRecursiveFibonacci_20-10  22.2µs ± 0%  22.2µs ± 0%   ~     (p=1.000 n=1+1)
SequentialFibonacci_10-10     3.73ns ± 0%  3.73ns ± 0%   ~     (p=1.000 n=1+1)
SequentialFibonacci_20-10     6.85ns ± 0%  6.84ns ± 0%   ~     (p=1.000 n=1+1)
SequentialFibonacci_40-10     13.1ns ± 0%  13.1ns ± 0%   ~     (p=1.000 n=1+1)
SequentialFibonacci_1000-10    320ns ± 0%   319ns ± 0%   ~     (p=1.000 n=1+1)
```

### References
- [Go Testing Package - Benchmarks](https://pkg.go.dev/testing#hdr-Benchmarks)
- [How to write benchmarks in go](https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go)
- [Go by Example: Testing and Benchmarking](https://gobyexample.com/testing-and-benchmarking)
- [Benchstat documentation](https://pkg.go.dev/golang.org/x/perf/cmd/benchstat#section-readme)
- [Overview of Benchmark Testing in Golang](https://www.geeksforgeeks.org/overview-of-benchmark-testing-in-golang/)