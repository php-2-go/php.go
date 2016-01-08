package php

import (
    _fmt "fmt"
    _str "strings"
)

const (
    VERSION = "1.0.0-beta"
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
// @param  i... interface{}
// @return (bool)
func IsEmpty(i... interface{}) (bool) {
    for _, iv := range i {
        if iv == nil || iv == false || iv == "" || iv == 0 || iv == 0.0 {
            return true
        }
    }
    return false
}
