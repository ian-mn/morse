package main

import (
	"fmt"
	"morse"
)

func main() {
	morse.SetSymbols("x", "o", "_", "-")
	morse.SetAlphabet("russian")

	encoded := morse.Encode("Привет, мир!")
	fmt.Println(encoded)

	decoded := morse.Decode(encoded)
	fmt.Println(decoded)
}