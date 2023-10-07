# morse
[![Go Report Card](https://goreportcard.com/badge/github.com/ian-mn/morse)](https://goreportcard.com/report/github.com/ian-mn/morse)
[![codecov](https://codecov.io/gh/ian-mn/morse/branch/master/graph/badge.svg)](https://codecov.io/gh/ian-mn/morse)
[![Go Reference](https://pkg.go.dev/badge/github.com/ian-mn/morse.svg)](https://pkg.go.dev/github.com/ian-mn/morse)
[![License](https://img.shields.io/pypi/l/Django.svg)](https://github.com/ian-mn/morse/blob/master/LICENSE)

Morse Code Library in Go with Extensible Alpahbets and Customization

## Download & Use
WIP

## Sample Usage
### Simple Encode/Decode
```go
package main

import (
	"fmt"
	"morse"
)

func main() {
	encoded := morse.Encode("CQD")
	fmt.Println(encoded) // -.-. --.- -..

	decoded := morse.Decode(encoded)
	fmt.Println(decoded) // CQD
}
```
### Adding Custom Alphabets
```go
package main

import (
	"fmt"
	"morse"
)

func main() {

	// Package is provided with international
	// and russian morse alphabets
	existingAlphabets := morse.GetAlphabets()
	fmt.Println(existingAlphabets) // [international russian]

	arabicAlpahbet := morse.Alphabet{
		Mapping: map[string]string{
			"ا": ".-",
			"ب": "-...",
			"ت": "-",
			"ث": "-.-.",
			"ج": ".---",
			"ح": "....",
			"خ": "---",
			"د": "---",
			"ذ": "-..",
			"ر": ".-.",
			"ز": "---.",
			"س": "...",
			"ش": "----",
			"ص": "-..-",
			"ض": "...-",
			"ط": "..-",
			"ظ": "-.--",
			"ع": ".-.-",
			"غ": "--.",
			"ف": "..-.",
			"ق": "--.-",
			"ك": "-.-",
			"ل": ".-..",
			"م": "--",
			"ن": "-.",
			"ه": "..-..",
			"و": ".--",
			"ي": "..",
			"ﺀ": ".",
		},
		Dot:  ".",
		Dash: "-",
	}

	morse.AddAlphabet(
		"arabic",
		arabicAlpahbet,
	)
	morse.SetAlphabet("arabic")

	encoded := morse.Encode("حبيبي") // My love
	fmt.Println(encoded) // .... -... .. -... ..
}
```
### Not Only Dots & Dashes
```go
package main

import (
	"fmt"
	"morse"
)

func main() {
	morse.SetSymbols("*", "^", "-", "\\")
	morse.SetAlphabet("russian")

	// Hello, World!
	encoded := morse.Encode("Привет, мир!")

	// *^^*-*^*-**-*^^-*-^-^^**^^\^^-**-*^*-^*^*^^
	fmt.Println(encoded)

	decoded := morse.Decode(encoded)
	fmt.Println(decoded) // ПРИВЕТ, МИР!
}
```
