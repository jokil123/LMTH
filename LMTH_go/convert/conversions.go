package convert

// html : lmth
var conh2l = map[string]string{
	"b": "d",
	"p": "q",
	"q": "p",
	"i": "!",
	"a": "ɐ", // ɒ
	"s": "ƨ",
	"u": "n", // υ
}

var conl2h map[string]string

func init() {
	conl2h = make(map[string]string)
	for k, v := range conh2l {
		conl2h[v] = k
	}
}
