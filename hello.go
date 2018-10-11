package main

import (
    "fmt"
    "runtime"
)

func main(){
    var b string = "runoob.com"
    var c bool
    a := runtime.Version()

    fmt.Println(a, b, c)
}