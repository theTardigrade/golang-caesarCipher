# golang-caesarCipher

This package provides simple functions to encrypt and decrypt text using a Caesar cipher.

## Blog Post

For more information, read [this blog post](https://golangprojectstructure.com/caesar-cipher-secret-messages/).

## Example

```golang
package main

import (
	"github.com/theTardigrade/golang-caesarCipher"

	"fmt"
)

func main() {
	const encryptShift = 15
	const input = "THIS IS A TEST"

	encrypted := caesarCipher.Encrypt(input, encryptShift)
	decrypted := caesarCipher.Decrypt(encrypted, encryptShift)

	fmt.Printf(
		"INPUT: %s\nENCRYPTED: %s\nDECRYPTED: %s\nINPUT EQUALS DECRYPTED: %t\n",
		input, encrypted, decrypted, input == decrypted,
	)
}
```