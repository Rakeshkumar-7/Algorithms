package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	words := []string{"hello", "hallow", "test", "smith", "smythe", "smote"}
	for _, word := range words {
		fmt.Printf("%s: %s\n", word, soundex(word))
	}
}

func soundex(str string) string {
	if len(str) == 0 {
		// edge case?
		return "A000"
	}

	// Source: https://en.wikipedia.org/wiki/Soundex
	articulationMap := map[byte]uint8{
		'a': 0, 'e': 0, 'i': 0, 'o': 0, 'u': 0, 'y': 0, 'h': 0, 'w': 0,
		'b': 1, 'f': 1, 'p': 1, 'v': 1,
		'c': 2, 'g': 2, 'j': 2, 'k': 2, 'q': 2, 's': 2, 'x': 2, 'z': 2,
		'd': 3, 't': 3,
		'l': 4,
		'm': 5, 'n': 5,
		'r': 6,
	}

	str = strings.ToLower(str)

	encoded := make([]uint8, 0)
	var prev uint8 = 0
	for i := 0; i < len(str); i++ {
		code, exists := articulationMap[str[i]]
		// if its not a alphabet or if it has the same code as the previous character, skip it
		if !exists || (i != 0 && code == prev) {
			continue
		}
		encoded = append(encoded, code)
		prev = code
	}

	var buffer strings.Builder
	// Write the upper-case equivalent of the first alphabetic character of our input
	for i := 0; i < len(str); i++ {
		if unicode.IsLetter(rune(str[i])) {
			buffer.WriteByte(byte(unicode.ToUpper(rune(str[i]))))
			break
		}
	}

	// add encoded value till the length of the output is 4 or the encoded values are exhausted
	for i := 1; i < len(encoded) && buffer.Len() < 4; i++ {
		// zeros are ignored - Read: https://web.stanford.edu/class/archive/cs/cs106b/cs106b.1216/assignments/1-cpp/soundex
		if encoded[i] == 0 {
			continue
		}
		buffer.WriteString(strconv.Itoa(int(encoded[i])))
	}

	// Padding with '0' if len(output) < 4
	for buffer.Len() < 4 {
		buffer.WriteString("0")
	}

	return buffer.String()
}
