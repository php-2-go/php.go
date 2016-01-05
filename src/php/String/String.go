package String

// func Len(s string) (uint) {}

// func Pos(s, ss string) (int, bool) {}
// func Posi(s, ss string) (int, bool) {}

func Rev(s string) (string) {
    sr := []rune(s)
    sl := len(sr)
    rr := make([]rune, sl)
    for i, ii := (sl - 1), 0; i >= 0; i, ii = (i - 1), (ii + 1) {
        rr[ii] = sr[i]
    }
    return string(rr)
}
