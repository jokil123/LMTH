package convert

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

func DecodeTag(s string) string {
	if v, ok := conl2h[s]; ok {
		return v
	}

	if len(s) > 1 {
		return ReverseString(s)
	}

	return s
}

var decRegexp = regexp.MustCompile(`<[\w\d]+>`)

// this is a recursive function that will decode the HTML
// it will find the first tag, decode it, then find the end tag
// then it will decode the content between the tags
func DecodeLMHT(s string) (string, error) {
	decStr := ""

	for {
		opLoc := decRegexp.FindStringIndex(s)
		if opLoc == nil {
			break
		}

		opStart := opLoc[0]
		opEnd := opLoc[1]

		tagName := s[opStart+1 : opEnd-1]

		// log.Printf("Found tag: %s\n", tagName)
		// log.Printf("finding closing tag for string: %s\n", s[opEnd:])
		clIndex := FindClosingTagIndex(s[opStart:], tagName) + opStart
		if clIndex == -1 {
			log.Println("Error decoding HTML: no end tag found")
			return "", nil
		}

		// decode the content between the tags
		content := s[opEnd:clIndex]
		decodedContent, err := DecodeLMHT(content)
		if err != nil {
			return "", err
		}

		// decodedTag := DecodeTag(tagName)
		decStr += fmt.Sprintf("<%s>%s</%s>", tagName, decodedContent, tagName)
		s = s[clIndex+len(tagName)+2:]
	}

	return decStr, nil
}

// find the string index of the closing tag (same level of nesting)
// assumes that the first tag is the opening tag
// tagName is without the <>
// example
// <html><lmth><html><lmth> -> 6
// <html><html><lmth><lmth> -> 18
func FindClosingTagIndex(s string, tagName string) int {
	// // log.Printf("Finding closing tag for string: %s\n", s)

	openingTag := "<" + tagName + ">"
	// log.Printf("Opening tag: %s\n", openingTag)
	closingTag := "<" + DecodeTag(tagName) + ">"
	// log.Printf("Closing tag: %s\n", closingTag)

	level := 0
	sI := 0

	for {
		// log.Printf("Searching string: %s\n", s[sI:])
		// log.Printf("search Index: %d\n", sI)

		opLoc := strings.Index(s[sI:], openingTag)
		clLoc := strings.Index(s[sI:], closingTag)

		// log.Printf("Opening tag index: %d\n", opLoc)
		// log.Printf("Closing tag index: %d\n", clLoc)

		nsi := 0

		if opLoc == -1 && clLoc == -1 {
			// log.Println("no tags found")
			return -1
		} else if opLoc == -1 {
			level--
			nsi = sI + clLoc + len(closingTag)
			// log.Println("no opening tag found")
		} else if clLoc == -1 {
			level++
			nsi = sI + clLoc + len(closingTag)
			// log.Println("no closing tag found")
		} else if clLoc > opLoc {
			level++
			nsi = sI + opLoc + len(openingTag)
			// log.Println("opening tag found first")
		} else if clLoc < opLoc {
			level--
			nsi = sI + clLoc + len(closingTag)
			// log.Println("closing tag found first")
		} else {
			// log.Println("error")
			return -1
		}

		// log.Printf("Level: %d\n", level)

		if level == 0 {
			// log.Printf("Found closing tag: %s\n", s[sI:nsi])
			return sI + clLoc
		}

		sI = nsi
	}
}

func DecodeFile(ip string, op string) error {
	f := LoadFile(ip)

	d, err := DecodeLMHT(f)
	if err != nil {
		log.Printf("Error decoding HTML: %s\n", err)
		return err
	}

	SaveFile(op, d)

	return nil
}
