# morse
Morse Code Library in Go

## Download and Use
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

	// Package is provided with international and russian morse alphabets
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
