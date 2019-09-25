package main

import "github.com/appliedgocourses/bank"

func main() {

    bank.Load()

    defer bank.Save()

}
