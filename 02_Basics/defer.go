package main

import (
    "fmt"
)

func trace(name string) func() {
    fmt.Println("Entering", name)
    // TODO:
    // 1. Print "Entering <name>"
    // 2. return a func() that prints "Leaving <name>" 
    return func() {
        fmt.Println("Leaving", name)
    }
}

func f() {
    defer trace("trace")()// TODO: add trace() here so the defer receives the returned function
    fmt.Println("Doing something")
}

func main() {
    fmt.Println("Before f")
    f()
    fmt.Println("After f")
}

