package main

import (
    "github.com/l3x/lazyserver/servers"
    "os"
    "fmt"
)

// To test:  curl http://localhost:1111/
func main() {
    lazyJob := &servers.LazyJob{Name: "lazy"}
    lazyFastJob := &servers.LazyFastJob{Name: "lazy_fast"}

    if len(os.Args) == 1 {
        fmt.Printf("Usage: lazyserver <lazy | lazy_fast>\nExample: lazyserver lazy\n")
        os.Exit(1)
    }
    if os.Args[1] == "lazy" {
        lazyJob.Run()
    } else {
        lazyFastJob.Run()
    }
}
