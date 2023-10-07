package main

import (
	"fmt"
	"morse"
)

func main() {
	encoded := morse.Encode("CQD")
	fmt.Println(encoded)

	decoded := morse.Decode(encoded)
	fmt.Println(decoded)
}
