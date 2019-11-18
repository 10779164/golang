package main

import (
    "fmt"
    "os"
)


func main() {
    var PWD string
    PWD = os.Getenv("PWD")
    fmt.Println(PWD)
}
