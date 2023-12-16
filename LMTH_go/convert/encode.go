package convert

import (
	"log"
	"regexp"
)

// div -> vid
func EncodeTag(s string) string {
	if v, ok := conh2l[s]; ok {
		return v
	}

	if len(s) > 1 {
		return ReverseString(s)
	}

	log.Printf("Unknown tag: %s\n", s)
	return s
}

var encRegexp = regexp.MustCompile(`<\/[\w\d]+>`)

func EncodeHTML(s string) (string, error) {
	for loc := encRegexp.FindStringIndex(s); loc != nil; loc = encRegexp.FindStringIndex(s) {
		start := loc[0]
		end := loc[1]

		tagName := s[start+2 : end-1]
		encodedTag := EncodeTag(tagName)

		s = s[:start+1] + encodedTag + s[end-1:]
	}

	return s, nil
}

func EncodeFile(ip string, op string) error {
	f := LoadFile(ip)

	e, err := EncodeHTML(f)
	if err != nil {
		log.Printf("Error encoding HTML: %s\n", err)
		return err
	}

	SaveFile(op, e)

	return nil
}
