```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        var count int
        var mutex sync.Mutex

        for i := 0; i < 1000; i++ {
                wg.Add(1)
                go func(i int) {
                        defer wg.Done()
                        mutex.Lock()
                        defer mutex.Unlock()
                        count += i
                }(i)
        }

        wg.Wait()
        fmt.Println(count) // Output will be 499500
}
```