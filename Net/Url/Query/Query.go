package Url

import (
    _fmt "fmt"
    _str "strings"
)

import (
    "php"
    php_net_url "php/Net/Url"
)

// Parse URL query.
//
// @param  q string
// @return (map[string]string)
func Parse(q string) (map[string]string) {
    r := make(map[string]string)
    // ensure explode action
    if _str.Index(q, "&") == -1 {
        q += "&"
    }
    if tmp := _str.Split(q, "&"); len(tmp) >= 2 {
        for _, tm := range tmp {
            if t := _str.SplitN(tm, "=", 2); len(t) >= 1 {
                var k, v string
                k = _str.TrimSpace(t[0])
                for i, _ := range t {
                    if i == 1 {
                        v = _str.TrimSpace(t[1])
                        break
                    }
                }
                r[k] = v
            }
        }
    }
    return r
}

// Unparse URL query.
//
// @param  q interface{}
// @return (string)
func Unparse(q interface{}) (string) {
    r := make([]string, 0)
    switch q := q.(type) {
        case map[string]string:
            for k, v := range q {
                r = append(r, _joinKeyValue(k, v))
            }
        case map[string]interface{}:
            for k, v := range q {
                r = append(r, _joinKeyValue(k, v))
            }
            break
    }
    return _str.Join(r, "&")
}

// Join key => value pairs (only 2-dims arrays for now).
//
// @param  k string
// @param  v interface{}
// @return (string)
func _joinKeyValue(k string, v interface{}) (string) {
    var s string
    switch v := v.(type) {
        case []int:
            for _, vv := range v {
                s += _fmt.Sprintf("%s[]=%v&", php_net_url.Encode(k), vv)
            }
            break
        case []string:
            for _, vv := range v {
                s += _fmt.Sprintf("%s[]=%v&", php_net_url.Encode(k), vv)
            }
            break
        case map[string]interface{}:
            for kk, vv := range v {
                switch vv := vv.(type) {
                    case []int, []string:
                        s += _joinKeyValue(k, vv)
                        break
                    case map[string]interface{}:
                        for kkk, vvv := range vv {
                            s += _joinKeyValue(kkk, vvv)
                        }
                        break
                    default:
                        s += _fmt.Sprintf("%s[%s]=%v&",
                            php_net_url.Encode(k),
                            php_net_url.Encode(kk),
                            php_net_url.Encode(php.String(vv),
                        ))
                }
            }
            break
        default:
            s = _fmt.Sprintf("%s=%v",
                php_net_url.Encode(k),
                php_net_url.Encode(php.String(v)),
            )
    }
    return _str.Trim(s, "&")
}
