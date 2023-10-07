package morse

import "testing"

func TestEncode(t *testing.T) {
	text := "Lorem ipsum dolor sit amet, consectetur adipiscing elit"
	correct := ".-.. --- .-. . --/.. .--. ... ..- --/-.. --- .-.. --- .-./... .. -/.- -- . - --..--/-.-. --- -. ... . -.-. - . - ..- .-./.- -.. .. .--. .. ... -.-. .. -. --./. .-.. .. -"
	got := Encode(text)
	if got != correct {
		t.Errorf(`Encoding %s:
		want: %s
		got: %s`, text, correct, got)
	}
}

func TestEncodeEmpty(t *testing.T) {
	text := ""
	correct := ""
	got := Encode(text)
	if got != correct {
		t.Errorf(`Encoding %s:
		want: %s
		got: %s`, text, correct, got)
	}
}

func TestEncodeNotInAlphabet(t *testing.T) {
	text := "Apple Jui~e"
	correct := ".- .--. .--. .-.. ./.--- ..- ..  ."
	got := Encode(text)
	if got != correct {
		t.Errorf(`Encoding %s:
		want: %s
		got: %s`, text, correct, got)
	}
}

func TestDecode(t *testing.T) {
	text := ".-.. --- .-. . --/.. .--. ... ..- --/-.. --- .-.. --- .-./... .. -/.- -- . - --..--/-.-. --- -. ... . -.-. - . - ..- .-./.- -.. .. .--. .. ... -.-. .. -. --./. .-.. .. -"
	correct := "LOREM IPSUM DOLOR SIT AMET, CONSECTETUR ADIPISCING ELIT"
	got := Decode(text)
	if got != correct {
		t.Errorf(`Encoding %s:
		want: %s
		got: %s`, text, correct, got)
	}
}

func TestDecodeEmpty(t *testing.T) {
	text := ""
	correct := ""
	got := Decode(text)
	if got != correct {
		t.Errorf(`Encoding %s:
		want: %s
		got: %s`, text, correct, got)
	}
}

func TestDecodeNotInAlphabet(t *testing.T) {
	text := ".- .--. .--. .-.. ./.--- ..- .. ~~ ."
	correct := "APPLE JUIE"
	got := Decode(text)
	if got != correct {
		t.Errorf(`Encoding %s:
		want: %s
		got: %s`, text, correct, got)
	}
}