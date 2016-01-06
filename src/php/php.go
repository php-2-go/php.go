package php

import (
    _fmt "fmt"
    _str "strings"
)

func Shutup() {}

// Get real type.
//
// @param  args.. interface{}
// @return (string)
func Type(args... interface{}) (string) {
    return _str.Replace(_fmt.Sprintf("%T", args[0]), " ", "", -1)
}

// Check is empty.
//
// @param  i interface{}
// @return (bool)
func Empty(i interface{}) (bool) {
    return (i == nil || i == "" || i == 0)
}
