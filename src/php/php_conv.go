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
    r, err := _strc.ParseInt(_fmt.Sprintf("%v", i), 10, 32)
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
    r, err := _strc.ParseFloat(_fmt.Sprintf("%v", i), 64)
    if err == nil {
        return float64(r)
    }
    return 0.00
}
