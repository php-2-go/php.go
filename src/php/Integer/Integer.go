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

// UInt converter.
//
// @param  i interface{}
// @return (uint)
func UInt(i interface{}) (uint) {
    if n := _n.Convert(i, "uint"); n != nil {
        return n.(uint)
    }
    return 0
}
