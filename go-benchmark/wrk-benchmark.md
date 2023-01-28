## Benchmarking HTTP endpoint using go wrk

wrk is a modern HTTP benchmarking tool capable of generating significant load when run on a single multi-core CPU

It will generate a significant load when run on a single multi-core CPU

### Installation
```
On Mac
brew install wrk

On Linux
sudo apt-get install build-essential libssl-dev git -y
git clone https://github.com/wg/wrk.git wrk
cd wrk
sudo make
# move the executable to somewhere in your PATH, ex:
sudo cp wrk /usr/local/bin

Using Docker
https://hub.docker.com/r/williamyeh/wrk/
```
### Usages
|Option|Description|
| ----------- | ----------- |
|-c, -connection|Amount of connections to keep open|
|-d, -duration|Duration of the benchmark process|
|-t, -threads|How many threads of process that will be invoked|

`wrk -t12 -c400 -d20s --latency http://localhost:8000/posts`

> This runs a benchmark for 20 seconds, using 12 threads in total, and keeping 400 HTTP connections open.
```
- c: the number of connections to use
- t: the number of threads to use
- d: the test duration, e.g., 60s
- s: the lua script to use for load testing our service (will cover in later section)
- timeout how many seconds to timeout a request
- latency: show the latency distribution for all the requests
```
For connections and threads, go wrk author suggests using thread number less than the core in CPU.
The connections are shared in different threads, *i.e.*, each threads get N = connections/threads connections.

### Configuration
What if we want to do a more intelligent benchmark, and be able to script to change some parameters?
it’s also possible with the parameter -s to send a lua script because wrk supports executing a LuaJIT script.
>Benchmarking POST requests using wrk

Example Lua script to configure a post request
Create a new .lua file containing the following code
```
wrk.method = "POST"
wrk.body = '{"title": "Harpot","description": "This is another harpot book","cover": "https://mantap.com/mantap.png","authorID": 1}'
wrk.headers["Content-Type"] = "application/json"
```
```
wrk -c1 -t1 -d1s -s ./benchmark/book-post.lua http://localhost:8000/books
```
It will execute the benchmark using POST method with the defined body and header content.

---

### Benchmarking Articles API
- Benchmark Get article list
  ```
  Without parameters
  wrk -t12 -c400 -d20s --latency http://localhost:8000/articles

  With parameters
  wrk -t12 -c400 -d20s --latency http://localhost:8000/articles\?limit\=100\&page\=2
  ```
- Benchmark Create one article
  ```
  wrk -c1 -t1 -d1s -s ./backend/benchmarks/create-article-post.lua http://localhost:8000/articles
  ```
  **Note:**`create-article-post.lua` is a LUA script that contains the configuration for wrk requests:
  ```
  wrk.method = "POST"
  wrk.body = '{"title": "Good morning~~","content": "Hello there!!","category_id": 5}'
  wrk.headers["Content-Type"] = "application/json"
  ```

### References
- [wrk benchmark benchmarking tool](https://github.com/wg/wrk)
- [Intelligent Benchmark with wrk](https://medium.com/@felipedutratine/intelligent-benchmark-with-wrk-163986c1587f)

