package main

import (
    "fmt"
    "strings"
)

func count(s string, m map[string]int){
    words := strings.Split(s, " ")
    for _, w := range words {
        w = strings.ToLower(w)
        w = strings.Trim(w, `\t\n\"'.,:;?!()-`)
        m[w]++
    }
    fmt.Println(m)
}

func main() {
    str := "foo bar baz! foo"
    strMap := map[string]int{}

    count(str, strMap)
}
