package php

import (
    _fmt "fmt"
    _str "strings"
)

func Shutup() {}

// Dump.
//
// @param  args... interface{}
// @return (void)
func Dump(args... interface{}) {
    _fmt.Println(args...)
}

// Dump as string.
//
// @param  args... interface{}
// @return (void)
func Dumps(args... interface{}) {
    var f string
    for _, arg := range args {
        _ = arg // silence..
        f += "%+v "
    }
    _fmt.Printf("%s\n", _fmt.Sprintf(f, args...))
}

// Dump types.
//
// @param  args... interface{}
// @return (void)
func Dumpt(args... interface{}) {
    var f string
    for _, arg := range args {
        _ = arg // silence..
        f += "%T "
    }
    _fmt.Printf("%s\n", _fmt.Sprintf(f, args...))
}

// Dump all.
//
// @param  args... interface{}
// @return (void)
func Dumpa(args... interface{}) {
    var f string
    for _, arg := range args {
        _ = arg // silence..
        f += "%#v "
    }
    _fmt.Printf("%s\n", _fmt.Sprintf(f, args...))
}

// Dump as formatted string.
//
// @param  f       string
// @param  args... interface{}
// @return (void)
func Dumpf(f string, args... interface{}) {
    _fmt.Printf("%s\n", _fmt.Sprintf(f, args...))
}

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
