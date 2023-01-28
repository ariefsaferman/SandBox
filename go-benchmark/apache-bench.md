## Benchmarking HTTP endpoint with Apache Benchmark

### Installation
As of macos Big Sur Apache Benchmark is installed by default in macos

### Installation using apt
```
sudo apt install apache2
```

### Usages
`ab -V`
Check apache benchmark version

`ab -n 10 http://localhost/`
Send request 10 times to the specified url

`ab -n 10 -c 10 -g run1.csv http://localhost/`
Output the result into a csv file as a tab separated values

`ab -n 10 -c 10 -g run1.csv http://localhost/`
Output the result into a csv file as a comma separated values

`ab -n 10 -c 10 -g run1.csv -v 2 http://localhost/`
How much troubleshooting info to print, the verbosity set to level 2

```
-n | requests, Number of requests to perform.
-c | concurrency, Number of multiple requests to perform at a time. Default is one request at a time.
-w | output in html format.
-v | How much troubleshooting info to print, set verbosity level - `4` max.
-t | time limit for benchmark process.
-g | output collected data to gnuplot format file.
-T | To set content type of the header (e.g: application/json).
-p | To specify the file containing data to post.
-H | adds an Auth header (could be Basic or Token)
```

---
### Examples

Example of benchmarking **POST** login request:
```
ab -n 1000 -c 200 -T application/json -p /benchmarks/login-post.json http://localhost:3000/v1/login
```

`login-post.json` file contains data that we want to use as the request body.
You can add more json file for other requests body.

---

Example of benchmarking **GET** request that requires an **bearer authentication token**
```
ab -n 1000 -c 500 -T application/json -H "Authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZW1haWwiOiJhY29uZ0BtYWlsLmNvbSIsImlzcyI6Imlzc3VlcmZvcmp3dCIsImV4cCI6MTY3MTEwNzg3OSwiaWF0IjoxNjcxMDg2Mjc5fQ.6qPvJJr4DvEV4fF8NDrnHUNs2WSeFK_P9u2sozz2wVg" -v 2 http://localhost:8000/details
```
---

Example of benchmarking **GET** request and put the result **into a csv file**

```
ab -n 1000 -c 100 -g current.txt http://localhost:3000/v1/trainers
```

Following is the result of the above command:
```
Server Software:        
Server Hostname:        localhost
Server Port:            3000

Document Path:          /v1/trainers
Document Length:        103 bytes

Concurrency Level:      100
Time taken for tests:   1.132 seconds
Complete requests:      1000
Failed requests:        0
Total transferred:      227000 bytes
HTML transferred:       103000 bytes
Requests per second:    883.28 [#/sec] (mean)
Time per request:       113.215 [ms] (mean)
Time per request:       1.132 [ms] (mean, across all concurrent requests)
Transfer rate:          195.81 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        1   24   8.9     25      75
Processing:     1   77  25.6     87     133
Waiting:        1   77  25.7     87     133
Total:          2  101  28.5    113     183

Percentage of the requests served within a certain time (ms)
  50%    113
  66%    115
  75%    116
  80%    117
  90%    125
  95%    134
  98%    147
  99%    160
 100%    183 (longest request)
```

1,132 seconds have been taken to complete the test and the longest response time is 183ms. Average response time is 113.215ms. Something interesting to note is the complete and failed requests.

- **Complete requests**: Number of successful responses received
- **Failed requests**: Number of requests that were considered a failure.

---
When doing a load tests or benchmarking, picking an arbitrary number and hitting your server is generally not a good way to do it. All the test resulr would tell is that your server can handle 100 or 200 concurent requests as long as they dont mind waiting for a spesific time for their request to load.
What you might want to do are:
1. Establish a baseline, use 1 visitor or 1 concurent process.
2. Second, start ramping up the numbers, e.g: 1, 10, 25, 50, 100, 125, etc..
3. Finally, make sure these requests run for prolonged periods of time, dont just start it and then stop it with ctrl+c in the terminal.

Once you have your results, graph them: number of visitors versus average request times, including max and min bars. Basically, load testing of an arbitrary application is only as useful as relevant tests; in this case, for example, if it takes 1 visitor 6s to load a page, then 7s a page for 200 visitors doesn't sound to bad, does it?

---

### Benchmarking Articles API
- Benchmark Get article list
  ```
  Without parameters
  ab -n 10000 -c 100 http://localhost:8000/articles

  With parameters
  ab -n 10000 -c 100 http://localhost:8000/articles\?limit\=100\&page\=2
  ```
- Benchmark Create one article
  ```
  ab -n 10000 -c 100 -p ./create-article-post.json http://localhost:8000/articles
  ```
  **Note:**`create-article-post.json` is a JSON file containing the body data of the request:
  ```
  {
    "title": "Good morning~~",
    "content": "Hello there!!",
    "category_id": 5
  }
  ```
---

**Notes:**
Running more than 1 benchmark test at the same time could cause the test to fail since it will overload the server and will result in failed requests.

### References
- [AB Documentation](https://httpd.apache.org/docs/2.4/programs/ab.html)
- [How to install apache benchmark](https://gist.github.com/yolossn/20d86c79745acbd97125b9cca950cbf7)
- [How to analyze apache bench result](https://serverfault.com/questions/378310/how-do-i-analyze-an-apache-bench-result)
- [Apache Bench - Environment Setup](https://www.tutorialspoint.com/apache_bench/apache_bench_environment_setup.htm)
- [Benchmark Menggunakan Apache Bench](https://medium.com/laravel-web-id/benchmark-menggunakan-apache-bench-7dee86716704)
- [Load testing a REST api using POST requests](https://erangad.medium.com/load-testing-a-rest-api-using-post-requests-6b0338196af0)