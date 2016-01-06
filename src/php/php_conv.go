package php

import (
    _fmt "fmt"
    _strc "strconv"
)

// Int converter.
//
// @param  i interface{}
// @return (int)
func Int(i interface{}) (int) {
    r, err := _strc.Atoi(_fmt.Sprintf("%v", i))
    if err == nil {
        return r
    }
    return 0
}
