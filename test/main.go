package main

import (
    "php"
    _u "php/Url"
    // _s "php/String"
)

func main() {
    x := _u.QueryParse("a=1&b")
    php.Dumpa(x)
    y := _u.QueryUnparse(x)
    php.Dumps(y)

    // s:= "0abcçdef"
    // php.Dump(_s.Pos(s, "ç"))
    // php.Dump(_s.Posi(s, "Ç"))
    // php.Dump(_s.Posr(s, "ç"))
    // php.Dump(_s.Posri(s, "Ç"))

    // php.Dump(_s.Len(s))
    // php.Dump(_s.Rev(s))

    // php.Dumpf("%#v", _s.Sub(s, -1, nil))
    // php.Dumpf("%#v", _s.Sub(s, -2, nil))
    // php.Dumpf("%#v", _s.Sub(s, -3, 1))
    // php.Dumpf("%#v", _s.Sub(s, 1, nil))
    // php.Dumpf("%#v", _s.Sub(s, 1, 1))
    // php.Dumpf("%#v", _s.Sub(s, 6, 2))
    // php.Dumpf("%#v", _s.Sub(s, 1, -6))
}
