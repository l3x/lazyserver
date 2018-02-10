package main

// Lazy and Very Random Server
import (
    "github.com/l3x/lazyserver.servers"
)

func main() {
    job := &servers.Job{
        Name:       "lazy",
    }

    job.Run()
}
