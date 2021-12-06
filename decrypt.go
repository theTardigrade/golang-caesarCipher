package caesarCipher

// Decrypt restores plaintext using an
// arbitrary-shift Caesar cipher.
func Decrypt(text string, shift uint) string {
	return Encrypt(text, 26-(shift%26))
}
