package main

import (
	"fmt"
	"morse"
)

func main() {

	existingAlphabets := morse.GetAlphabets()
	fmt.Println(existingAlphabets)

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
	fmt.Println(encoded)
}
