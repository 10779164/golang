package main

import(
    "fmt"
    "time"
)

func test(a string) {
    fmt.Println(a)
    return
}

func main() {
    var a string = "hello"
    fmt.Println("hello gaoyao!~")
    b := a+" world"
    fmt.Println(b)
    time.Sleep(time.Duration(3)*time.Second)
    go test("hello gx")
    sql := fmt.Sprintf("insert into %s value(1,2,3)",b)
    fmt.Println(sql)
}
