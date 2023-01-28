# Go - Profiling

## Why do we need profiling?
Profiling are used by developers to help identify performance problems without having to touch their code.

Profilers can answer questions like, **“How many times is each method in my code called?”** and, **“How long do each of these methods take?”** 

Profilers also track things like **memory allocations** and **garbage collection**. Some profilers can even track key methods in your code, so you can understand how often SQL statements and web services are called. In addition, some profilers can track web requests and train those transactions to understand the performance of transactions within your code.

Code profilers can track all the way down to each individual line of code. However, most developers only use profilers when chasing a CPU or memory problem, and need to go out of their way to try and find those problems. This is because many profilers make applications run a hundred times slower than usual. While most consider profilers to be a situational tool not meant for daily use, code profiling can be a total lifesaver when you need it.

## How to profile your Go Application
If you are using Gin, you will need to use the `gin-contrib/pprof`
```
https://github.com/gin-contrib/pprof
```

First you will need to register the pprof api in your router
```
package main

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()
  pprof.Register(router)
  router.Run(":8080")
}
```

by default the route for the pprof will be `/debug/pprof`. But you can change it by manually asign a new route

```
func main() {
	router := gin.Default()
	// default is "debug/pprof"
	pprof.Register(router, "dev/pprof")
	router.Run(":8080")
}
```

After you set up your profiler, run it using
```
//to get the memory allocation
go tool pprof http://localhost:8080/debug/pprof/heap

//to get the cpu usage
go tool pprof http://localhost:8080/debug/pprof/profile
```

If you want to have an HTML view of the profile,run it using
```
go tool pprof -http localhost:3435 app_url
```

### References
- [Profiling with pprof](https://pkg.go.dev/net/http/pprof)
- [Go-pprof error - graphvix](https://stackoverflow.com/questions/26216628/go-pprof-not-working-properly)
- [Performance Measuring, Profiling, and Optimizing Tips for Go Web Applications](https://articles.wesionary.team/performance-measuring-profiling-and-optimizing-tips-for-go-web-applications-20f2f812ff6e)

