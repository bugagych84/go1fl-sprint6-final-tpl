package service

import (
	"errors"
	"strings"
	"unicode"

	"github.com/bugagych84/go1fl-sprint6-final-tpl/pkg/morse"
)

var ErrEmpty = errors.New("empty input")

func DetectAndConvert(in string) (string, bool, error) {
	trimmed := strings.TrimSpace(in)
	if trimmed == "" {
		return "", false, ErrEmpty
	}

	hasAlnum := strings.IndexFunc(trimmed, func(r rune) bool {
		return unicode.IsLetter(r) || unicode.IsDigit(r)
	}) >= 0

	if hasAlnum {
		normalized := normalizeSpaces(trimmed)
		return morse.ToMorse(normalized), false, nil
	}

	isOnlyMorseAlphabet := strings.IndexFunc(trimmed, func(r rune) bool {
		switch r {
		case '.', '-', ' ', '/', '(', ')', '\'', '"', '=', '+', ',', '?', ':':
			return false
		default:
			if unicode.IsSpace(r) {
				return false
			}
			return true
		}
	}) == -1

	if isOnlyMorseAlphabet {
		fields := strings.Fields(trimmed)
		reconstructed := strings.Join(fields, " ")
		return morse.ToText(reconstructed), true, nil
	}

	normalized := normalizeSpaces(trimmed)
	return morse.ToMorse(normalized), false, nil
}

func normalizeSpaces(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return ' '
		}
		return r
	}, s)
}
