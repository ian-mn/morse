package morse

import (
	"strings"
)

func Encode(text string) string {
	upperText := strings.ToUpper(text)

	words := strings.Split(upperText, " ")
	var encoded string

	for i, word := range words {
		elems := strings.Split(word, "")
		for j, elem := range elems {
			encoded += config.encodingAlphabet.Mapping[elem]
			if j != len(elems)-1 {
				encoded += config.separator
			}
		}
		if i != len(words)-1 {
			encoded += config.space
		}
	}
	return encoded
}

func Decode(text string) string {
	words := strings.Split(text, config.space)
	var decoded string
	for i, word := range words {
		elems := strings.Split(word, config.separator)
		for _, elem := range elems {
			decoded += config.decodingAlphabet.Mapping[elem]
		}
		if i != len(words)-1 {
			decoded += " "
		}

	}
	return decoded
}
