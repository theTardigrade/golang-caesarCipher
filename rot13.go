package caesarCipher

// Rot13 both encrypts and decrypts text
// using a 13-shift Caesar cipher.
func Rot13(text string) string {
	return Encrypt(text, 13)
}
