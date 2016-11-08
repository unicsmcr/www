package crypto

import (
	r "github.com/hacksoc-manchester/www/helpers/rand"
	"testing"
)

func TestCrypto(t *testing.T) {
	SetSymmetricKey(r.RandString(32))
	var decryptedValue string
	for i := 0; i < 10; i++ {
		value := r.RandString(int(r.Src().Int63() % 101))
		encryptedValue, err := Encrypt(value)
		if err != nil {
			panic(err)
		}
		decryptedValue, err = Decrypt(encryptedValue)
		if err != nil {
			panic(err)
		} else if value != decryptedValue {
			t.Error(
				"For", value,
				"expected result", value,
				"got", decryptedValue,
			)
		}
	}
}
