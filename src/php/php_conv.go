package php

import (
    _fmt "fmt"
    _strc "strconv"
)

// Int converter.
//
// @param  i interface{}
// @return (int64)
func Int(i interface{}) (int64) {
    r, err := _strc.ParseInt(String(i), 10, 32)
    if err == nil {
        return (r)
    }
    return 0
}

// Float converter.
//
// @param  i interface{}
// @return (float64)
func Float(i interface{}) (float64) {
    r, err := _strc.ParseFloat(String(i), 64)
    if err == nil {
        return float64(r)
    }
    return 0.00
}

// String converter.
//
// @param  i interface{}
// @return (string)
func String(i interface{}) (string) {
    return _fmt.Sprintf("%v", i)
}
