package morse

import (
	"errors"
	"strings"
)

type Alphabet struct {
	Mapping map[string]string
	Dot     string
	Dash    string
}

var alphabets = map[string]Alphabet{

	"international": {
		Mapping: map[string]string{
			"A":  ".-",
			"B":  "-...",
			"C":  "-.-.",
			"D":  "-..",
			"E":  ".",
			"F":  "..-.",
			"G":  "--.",
			"H":  "....",
			"I":  "..",
			"J":  ".---",
			"K":  "-.-",
			"L":  ".-..",
			"M":  "--",
			"N":  "-.",
			"O":  "---",
			"P":  ".--.",
			"Q":  "--.-",
			"R":  ".-.",
			"S":  "...",
			"T":  "-",
			"U":  "..-",
			"V":  "...-",
			"W":  ".--",
			"X":  "-..-",
			"Y":  "-.--",
			"Z":  "--..",
			"1":  ".----",
			"2":  "..---",
			"3":  "...--",
			"4":  "....-",
			"5":  ".....",
			"6":  "-....",
			"7":  "--...",
			"8":  "---..",
			"9":  "----.",
			"0":  "-----",
			".":  ".-.-.-",
			":":  "---...",
			",":  "--..--",
			";":  "-.-.-",
			"?":  "..--..",
			"=":  "-...-",
			"'":  ".----.",
			"!":  "-.-.--",
			"-":  "-....-",
			"_":  "..--.-",
			"\"": ".-..-.",
			"(":  "-.--.",
			")":  "-.--.-",
			"$":  "...-..-",
			"&":  ".-...",
			"@":  ".--.-.",
			"+":  ".-.-.",
		},
		Dot:  ".",
		Dash: "-",
	},

	"russian": {
		Mapping: map[string]string{
			"А":  ".-",
			"Г":  "--.",
			"Ж":  "...-",
			"Й":  ".---",
			"М":  "--",
			"П":  ".--.",
			"Т":  "-",
			"Х":  "....",
			"Ш":  "----",
			"Ы":  "-.--",
			"Ю":  "..--",
			"Є":  "..-..",
			"Б":  "-...",
			"Д":  "-..",
			"З":  "--..",
			"К":  "-.-",
			"Н":  "-.",
			"Р":  ".-.",
			"У":  "..-",
			"Ц":  "-.-.",
			"Щ":  "--.-",
			"Ь":  "-..-",
			"Я":  ".-.-",
			"В":  ".--",
			"Е":  ".",
			"И":  "..",
			"Л":  ".-..",
			"О":  "---",
			"С":  "...",
			"Ф":  "..-.",
			"Ч":  "---.",
			"Ъ":  "--.--",
			"Э":  "..-..",
			"Ї":  ".---.",
			"Ґ":  "--.",
			"0":  "-----",
			"3":  "...--",
			"6":  "-....",
			"9":  "----.",
			"1":  ".----",
			"4":  "....-",
			"7":  "--...",
			"2":  "..---",
			"5":  ".....",
			"8":  "---..",
			".":  ".-.-.-",
			"'":  ".----.",
			"(":  "-.--.",
			":":  "---...",
			"+":  ".-.-.",
			"\"": ".-..-.",
			"¿":  "..-.-",
			",":  "--..--",
			"!":  "-.-.--",
			")":  "-.--.-",
			";":  "-.-.-.",
			"-":  "-....-",
			"$":  "...-..-",
			"¡":  "--...-",
			"?":  "..--..",
			"/":  "-..-.",
			"&":  ".-...",
			"=":  "-...-",
			"_":  "..--.-",
			"@":  ".--.-.",
		},
		Dot:  ".",
		Dash: "-",
	},
}

func (alpha Alphabet) swapKeyValue() Alphabet {
	swapped := map[string]string{}
	for k, v := range alpha.Mapping {
		swapped[v] = k
	}
	return Alphabet{Mapping: swapped, Dot: alpha.Dot, Dash: alpha.Dash}
}

func AddAlphabet(name string, alpha Alphabet) error {

	if len(alpha.Dot) != 1 || len(alpha.Dash) != 1 {
		return errors.New("dot and dash each should be one character")
	}

	if alpha.Dot == alpha.Dash {
		return errors.New("dot and dash should be different")
	}

	upperCaseMapping := map[string]string{}

	for k, v := range alpha.Mapping {

		for _, symbol := range strings.Split(v, "") {
			if symbol != alpha.Dot && symbol != alpha.Dash {
				return errors.New("unexpected symbol")
			}
		}

		upperCaseMapping[strings.ToUpper(k)] = strings.ToUpper(v)
	}

	alphabets[name] = Alphabet{
		Mapping: upperCaseMapping,
		Dot:     alpha.Dot,
		Dash:    alpha.Dash,
	}

	return nil
}

func GetAlphabets() []string {
	var alphaList = []string{}

	for k := range alphabets {
		alphaList = append(alphaList, k)
	}
	return alphaList
}
