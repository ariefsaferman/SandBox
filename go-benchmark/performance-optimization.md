# Performance Optimization
Optimization is really essential in building scalable application, to know where or how much optimization that we need, we have to do **benchmark and profilling** of our application to know the performance of our application before and after we do the optimization and to know in which part spesifically we need to do the optimization.

### Why we need to Optimize our system?
- Know what our system's needs are
- Reduce any unused resources and processes
- Increase our API performance

### Steps to optimize REST API
1. Have a real application that ready to be deployed or used
2. Do benchmark
3. Decide if the performance are above the standards (based on teams) or not, if not do profilling.
4. Get the result data from profiling to know which part of the code that cause the performance problem.
5. Decision, optimize the code.
6. Repeat to #2 until we are confident to deploy the API or our API have been improved significantly.
7. Deploy

### Tools for to help optimize our system
HTTP Load testing / Benchmarking tools
- [AB (Apache Benchmark)](https://httpd.apache.org/docs/2.4/programs/ab.html)
- [Wrk](https://github.com/wg/wrk)
- [Hey](https://github.com/rakyll/hey)
- [K6](https://github.com/grafana/k6?ref=thechiefio)
- Or any tools that can do many requests at once

Profiling tools
- [pprof](https://pkg.go.dev/net/http/pprof)
