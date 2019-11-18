package main

import "fmt"
import "time"

func test(a string) {
    for {
    for i := 0; i < 3; i++{
        fmt.Println(i)
	time.Sleep(time.Second)
	
    }
    time.Sleep(time.Second)  
    }
}

func main() {
     test("hello")
}
