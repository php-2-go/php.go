package main

import (
    "fmt"
    _s "php/String"
)

func main() {
    s:= "abcç"
    fmt.Println(_s.Rev(s))
}
