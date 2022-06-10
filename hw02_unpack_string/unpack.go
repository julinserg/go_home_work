package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func buildOneSymbol(prevSymbol rune, prevSymbolExist bool, strBuilder *strings.Builder) {
	if prevSymbolExist {
		strBuilder.WriteRune(prevSymbol)
	}
}

func buildSequenceSymbols(symbol rune, prevSymbol rune, prevSymbolExist bool, strBuilder *strings.Builder) error {
	if !prevSymbolExist {
		return ErrInvalidString
	}

	digit, err := strconv.Atoi(string(symbol))
	if err != nil {
		return err
	}

	if digit != 0 {
		for i := 0; i < digit; i++ {
			strBuilder.WriteRune(prevSymbol)
		}
	}
	return nil
}

func Unpack(str string) (string, error) {
	if len(str) == 0 {
		return "", nil
	}
	var resultStrBuilder strings.Builder
	var prevSymbol rune
	prevSymbolExist := false
	for _, symbol := range str {
		if unicode.IsDigit(symbol) {
			err := buildSequenceSymbols(symbol, prevSymbol, prevSymbolExist, &resultStrBuilder)
			if err != nil {
				return "", err
			}
			prevSymbolExist = false
		} else {
			buildOneSymbol(prevSymbol, prevSymbolExist, &resultStrBuilder)
			prevSymbol = symbol
			prevSymbolExist = true
		}
	}
	if prevSymbolExist {
		resultStrBuilder.WriteRune(prevSymbol)
	}

	return resultStrBuilder.String(), nil
}
