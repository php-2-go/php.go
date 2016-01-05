package String

import (
    _str "strings"
)

import (
    _i "php/Integer"
)

func Len(s string) (int) {
    return len([]rune(s))
}

func Pos(s, ss string) (int) {
    return _str.Index(s, ss)
}

// func Posi(s, ss string) (int) {}

func Rev(s string) (string) {
    sr := []rune(s)
    sl := len(sr)
    rr := make([]rune, sl)
    for i, ii := (sl - 1), 0; i >= 0; i, ii = (i - 1), (ii + 1) {
        rr[ii] = sr[i]
    }
    return string(rr)
}

func Sub(s string, ss, sl interface{}) (string) {
    rs := []rune(s)
    rl := len(rs)
    ssi, sli := _i.Int(ss), _i.Int(sl)
    ssl, sll := ssi, sli
    // make abs
    if ssi <= -1 { ssl = (ssi * -1) }
    if sli <= -1 { sll = (sli * -1) }
    // out of range returns ""
    if ssl >= rl { return "" }
    if sll >= rl { return "" }
    if sl == nil {
        if ssi <= -1 {
            rr := []rune(Rev(s))
            if ssl >= len(rr) {
                return s
            }
            var rs string
            for i := 0; i < ssl; i++ {
                rs = string(rr[i]) + rs
            }
            return rs
        }
        return string(rs[ssl:])
    }
    if sli >= 1 {
        // prevent "out of range" error
        if sll > rl {
            return string(rs[ssl: rl])
        }
        return string(rs[ssl: ssl + sll])
    } else {
        return string(rs[ssl: len(rs) - sll])
    }
}
