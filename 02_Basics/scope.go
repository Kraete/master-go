package main

import (
    "fmt"
    "unicode"
)

func main() {
    s := "abcde"
    for _, letter := range s {
        up := unicode.ToUpper(letter)
        fmt.Print(string(up))
    }
    fmt.Println("\n" + s)
}
