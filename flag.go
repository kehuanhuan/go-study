package main

import (
    "flag"
    "fmt"
)

func main() {
    var name string

    flag.StringVar(&name, "name", "test", "this is a test")

    flag.Parse()

    fmt.Println(name)

}

