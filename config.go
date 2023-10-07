package morse

import (
	"errors"
	"strings"
)

var (
	defaultDish      = "."
	defaultDash      = "-"
	defaultSeparator = " "
	defaultSpace     = "/"
	defaultAlphabet  = "international"
)

type cfg struct {
	dot              string
	dash             string
	separator        string
	space            string
	encodingAlphabet Alphabet
	decodingAlphabet Alphabet
}

var config = cfg{
	dot:              defaultDish,
	dash:             defaultDash,
	separator:        defaultSeparator,
	space:            defaultSpace,
	encodingAlphabet: alphabets[defaultAlphabet],
	decodingAlphabet: alphabets[defaultAlphabet].swapKeyValue(),
}

func SetSymbols(dot, dash, sep, space string) error {

	if len(dot) != 1 || len(dash) != 1 || len(sep) != 1 || len(space) != 1 {
		return errors.New("dot, dash, separator and space each should be one character")
	}

	for k, v := range config.encodingAlphabet.Mapping {
		newV := strings.Replace(v, config.encodingAlphabet.Dot, dot, -1)
		newV = strings.Replace(newV, config.encodingAlphabet.Dash, dash, -1)
		config.encodingAlphabet.Mapping[k] = newV
	}

	config.encodingAlphabet.Dot = dot
	config.encodingAlphabet.Dash = dash
	config.decodingAlphabet = config.encodingAlphabet.swapKeyValue()

	config.dot = dot
	config.dash = dash
	config.separator = sep
	config.space = space

	return nil
}

func SetAlphabet(key string) error {

	a, ok := alphabets[key]
	if ok {
		config.encodingAlphabet = a
	} else {
		return errors.New("such alphabet doesn't exist")
	}

	SetSymbols(config.dot, config.dash, config.separator, config.space)

	return nil
}
