package convert

import (
	"log"
	"os"
)

func ReverseString(s string) string {
	// Convert string to a slice of runes
	runes := []rune(s)

	// Get the length of the slice
	length := len(runes)

	// Reverse the order of the runes in the slice
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// Convert the slice of runes back to a string
	reversedString := string(runes)

	return reversedString
}

func SaveFile(p string, s string) {
	f, err := os.Create(p)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()

	_, err = f.WriteString(s)
	if err != nil {
		log.Println(err)
		return
	}
}

func LoadFile(p string) string {
	bytes, err := os.ReadFile(p)
	if err != nil {
		log.Println(err)
		return ""
	}
	return string(bytes)
}
