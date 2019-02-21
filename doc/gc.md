## Points
- Go provides excellent profiling tools that can point directly to the allocation-heavy portions of a code base?
- Go allocates memory in two places: a global heap for dynamic allocations and a local stack for each goroutine.
- stack allocation is cheap and heap allocation is expensive
- Stack allocation requires that the lifetime and memory footprint of a variable can be determined at compile time
- escape analysis : by the variable's scope
```
package main

import "fmt"

func main() {
        x := 42
        fmt.Println(x)
}
```

```
go build -gcflags '-m' ./main.go
# command-line-arguments
./main.go:7: x escapes to heap
./main.go:7: main ... argument does not escape, dynamically allocated on the heap at runtime
```

-
## articles
- https://blog.golang.org/profiling-go-programs
- https://segment.com/blog/allocation-efficiency-in-high-performance-go-services/
- https://en.wikipedia.org/wiki/Memory_management#HEAP
- https://en.wikipedia.org/wiki/Escape_analysis
-
