package main

import (
    "fmt"
    "os"
    "strings"
    "unicode"
)

func acronym(s string) (acr string) {
    for _, char := range s {
        if unicode.IsUpper(char) {
            acr = acr + string(char)
        }
    }
    return acr
}

func main() {
    s := "Pan Galactic Gargle Blaster" 
    if len(os.Args) > 1 {
        s = strings.Join(os.Args, " ")
    }
    fmt.Println(acronym(s))
}
