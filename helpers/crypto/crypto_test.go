package crypto

import (
	h "github.com/hacksoc-manchester/www/helpers"
	"testing"
)

func TestCrypto(t *testing.T) {
	SetSymmetricKey(h.RandString(32))
	for i := 0; i < 10; i++ {
		val := h.RandString(int(h.Src().Int63() % 101))
		res, err := Encrypt(val)
		if err != nil {
			panic(err)
		}
		res, err = Decrypt(res)
		if err != nil {
			panic(err)
		} else if val != res {
			t.Error(
				"For ", val,
				"expected result ", val,
				"got ", res,
			)
		}
	}
}
