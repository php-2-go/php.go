package Number

import (
    _fmt "fmt"
    _strc "strconv"
)

// Converter.
//
// @param  i  interface{}
// @param  it string
// @return (interface{})
func Convert(i interface{}, it string) (interface{}) {
    if i != nil {
        n, err := _strc.Atoi(_fmt.Sprintf("%v", i))
        if err == nil {
            switch it {
                // signed
                case    "int": return int(n)
                case   "int8": return int8(n)
                case  "int16": return int16(n)
                case  "int32": return int32(n)
                case  "int64": return int64(n)
                // unsigned
                case   "uint": return uint(n)
                case  "uint8": return uint8(n)
                case "uint16": return uint16(n)
                case "uint32": return uint32(n)
                case "uint64": return uint64(n)
                // float
                case "float32": return float32(n)
                case "float64": return float64(n)
            }
        }
    }
    return nil
}
