package Url

import (
    _url "net/url"
)

// URL encode.
//
// @param  s string
// @return (string)
func Encode(s string) (string) {
    return _url.QueryEscape(s)
}

// URL decode.
//
// @param  s string
// @return (string)
func Decode(s string) (string) {
    if s, err := _url.QueryUnescape(s); err == nil {
        return s
    }
    return ""
}
