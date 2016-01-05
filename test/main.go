package main

import (
    "fmt"
    _s "php/String"
)

// println()

func main() {
    s:= "0abcçdef"
    fmt.Println(_s.Pos(s, "ç"))
    fmt.Println(_s.Posi(s, "Ç"))
    fmt.Println(_s.Posr(s, "ç"))
    fmt.Println(_s.Posri(s, "Ç"))

    // fmt.Println(_s.Len(s))
    // fmt.Println(_s.Rev(s))

    // fmt.Printf("%#v\n", _s.Sub(s, -1, nil))
    // fmt.Printf("%#v\n", _s.Sub(s, -2, nil))
    // fmt.Printf("%#v\n", _s.Sub(s, -3, 1))
    // fmt.Printf("%#v\n", _s.Sub(s, 1, nil))
    // fmt.Printf("%#v\n", _s.Sub(s, 1, 1))
    // fmt.Printf("%#v\n", _s.Sub(s, 6, 2))
    // fmt.Printf("%#v\n", _s.Sub(s, 1, -6))
}
