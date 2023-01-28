## Question Bank

### Benchmark / Load Testing

1. Mention 3 Performance optimization tools that is used for benchmarking / load testing an http server
```
- apache bench
- hey
- K6 grafana
- wrk
```

2. Write down the command to run a benchmark using go testing library and specify the N times to run to be 100 times.
```
go test -bench=. benchtime=100x
```

3. What are benchstat used for in benchmarking with go testing library
```
Benchstat computes and compares statistics about benchmarks
```

4. When benchmarking an http requests using apache bench, what is the file type that contains the body of the requests?
```
.json file
```

5. When benchmarking using wrk, we can configure the requests that are being benchmarked, what are the file type used for the script?
```
LUA script
```

---

### Profiling
1. Why do we need profiling?
```
Profiling are used by developers to help identify performance problems without having to touch their code.
```
2. What is the package that you need to get if you want to do profiling and you are using gin
```
`gin-contrib/pprof`
```
3.
4.
5.
