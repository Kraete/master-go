package main

import (
    "fmt"
)

func longest(strings ...string) int {
    var max int
    for _, str := range(strings) {
        if len(str) > max {
            max = len(str)
        }
    }

    return max
}

func main() {
    fmt.Println(longest("Six", "sleek", "swans", "swam", "swiftly", "southwards"))
}
