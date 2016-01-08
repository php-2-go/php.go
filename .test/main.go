package main

import (
    "php"
    // _s "php/String"
    php_net_url "php/Net/Url"
    php_net_url_query "php/Net/Url/Query"
)

func main() {
    // x := "1.1"
    // php.Dumpa(php.Int(x))
    // php.Dumpa(php.Float(x))
    // php.Dumpa(php.Bool(x))
    // php.Dumpa(php.String(x))

    php.Dump(php_net_url.Encode("a=1"))
    php.Dump(php_net_url.Decode(php_net_url.Encode("a=1")))
    php.Dump(php_net_url_query.Parse("a=1"))
    php.Dump(php_net_url_query.Unparse(php_net_url_query.Parse("a=1")))

    // s:= "0açe"
    // php.Dump(_s.Index(s, "ç"))
    // php.Dump(_s.Index(s, "ç"))
    // php.Dump(_s.IndexNC(s, "Ç"))
    // php.Dump(_s.IndexLast(s, "ç"))
    // php.Dump(_s.IndexLastNC(s, "Ç"))

    // php.Dump(_s.Len(s))
    // php.Dump(_s.Reverse(s))

    // php.Dump(_s.Ord("♥"))
    // php.Dump(_s.Chr(_s.Ord("♥")))

    // php.Dumpf("%#v", _s.Substring(s, -1, nil))
    // php.Dumpf("%#v", _s.Substring(s, -2, nil))
    // php.Dumpf("%#v", _s.Substring(s, -3, 1))
    // php.Dumpf("%#v", _s.Substring(s, 1, nil))
    // php.Dumpf("%#v", _s.Substring(s, 1, 1))
    // php.Dumpf("%#v", _s.Substring(s, 6, 2))
    // php.Dumpf("%#v", _s.Substring(s, 1, -6))
}
