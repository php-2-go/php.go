package php

import (
    _fmt "fmt"
    _rex "regexp"
    _strc "strconv"
)

// Int converter.
//
// @param  i interface{}
// @return (int)
func Int(i interface{}) (int) {
    if r, err := _strc.ParseInt(String(i), 10, 64); err == nil {
        return int(r)
    }
    return 0
}

// Float converter.
//
// @param  i interface{}
// @return (float64)
func Float(i interface{}) (float64) {
    if r, err := _strc.ParseFloat(String(i), 64); err == nil {
        return r
    }
    return 0.0
}

// Bool converter.
//
// @param  i interface{}
// @return (float64)
func Bool(i interface{}) (bool) {
    if i == "0" || Empty(i) {
        return false
    }
    return true
}

// String converter.
//
// @param  i interface{}
// @return (string)
// @panics
func String(i interface{}) (string) {
    switch i.(type) {
        case nil:
            return ""
        case int, bool, string:
            return _fmt.Sprintf("%v", i)
        default:
            it := Type(i)
            // check numerics
            if ok, _ := _rex.MatchString("u?int(\\d+)?|float(32|64)", it); ok {
                return _fmt.Sprintf("%v", i)
            }
            panic("Unsupported input type '"+ it +"' given!")
    }
}