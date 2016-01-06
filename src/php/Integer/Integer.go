package Number

import (
    _n "php/Number"
)

// Int converter.
//
// @param  i interface{}
// @return (int)
func Int(i interface{}) (int) {
    if n := _n.Convert(i, "int"); n != nil {
        return n.(int)
    }
    return 0
}
