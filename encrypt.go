package caesarCipher

import (
	"strings"
)

func encryptShiftRune(r rune, shift uint) rune {
	var min, max rune

	if r >= 'A' && r <= 'Z' {
		min, max = 'A', 'Z'
	} else if r >= 'a' && r <= 'z' {
		min, max = 'a', 'z'
	}

	if max > min {
		d := (max - min) + 1
		s := rune(shift) % d

		if s == 0 {
			return r
		}

		return (((r - min) + s) % d) + min
	}

	return r
}

// Encrypt obscures plaintext using an
// arbitrary-shift Caesar cipher.
func Encrypt(text string, shift uint) string {
	var builder strings.Builder

	for _, r := range text {
		builder.WriteRune(encryptShiftRune(r, shift))
	}

	return builder.String()
}
