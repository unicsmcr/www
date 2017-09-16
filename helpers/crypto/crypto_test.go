package crypto

import (
	"testing"

	"github.com/hacksoc-manchester/www/helpers/rand"
)

func TestCrypto(t *testing.T) {
	SetSymmetricKey(rand.String(32))
	for i := 0; i < 10; i++ {
		value := rand.String(int(rand.Src().Int63() % 101))
		encryptedValue, err := Encrypt(value)
		if err != nil {
			panic(err)
		}

		decryptedValue, err := Decrypt(encryptedValue)
		if err != nil {
			panic(err)
		}

		if value != decryptedValue {
			t.Error(
				"For", value,
				"expected result", value,
				"got", decryptedValue,
			)
		}
	}
}
