package crypto

import (
	"github.com/hacksoc-manchester/www/helpers/rand"
	"testing"
)

func TestCrypto(t *testing.T) {
	SetSymmetricKey(rand.RandString(32))
	var decryptedValue string
	for i := 0; i < 10; i++ {
		value := rand.RandString(int(rand.Src().Int63() % 101))
		encryptedValue, err := Encrypt(value)
		if err != nil {
			panic(err)
		}

		decryptedValue, err = Decrypt(encryptedValue)
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
