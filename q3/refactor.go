package main

import (
	"fmt"
	"strings"
)

func findFirstStringInBracket(str string) string {
	if len(str) <= 0 {
		return ""
	}

	indexFirstBracket := strings.Index(str, "(")
	if indexFirstBracket < 0 {
		return ""
	}

	wordsAfterFirstBracket := str[indexFirstBracket+1:]
	indexClosingBracketFound := strings.Index(wordsAfterFirstBracket, ")")
	if indexClosingBracketFound <= 0 {
		return ""
	}

	return wordsAfterFirstBracket[:indexClosingBracketFound]
}

func main() {
	s := "Bast(ia)n)"
	fmt.Println(findFirstStringInBracket(s))
}

