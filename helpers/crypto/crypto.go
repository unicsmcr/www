package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"os"
)

var block cipher.Block

func init() {
	var err error

	block, err = aes.NewCipher([]byte(os.Getenv("SYMMETRIC_KEY")))

	if err != nil {
		panic(err)
	}
}

// Encrypt encrypts the specified text using AES. Returns a base64 encoded string.
func Encrypt(value string) (string, error) {
	base64Encoding := base64.URLEncoding.EncodeToString([]byte(value))
	encryptedValue := make([]byte, aes.BlockSize+len(base64Encoding))
	iv := encryptedValue[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(encryptedValue[aes.BlockSize:], []byte(base64Encoding))

	return base64.URLEncoding.EncodeToString(encryptedValue), nil
}

// Decrypt decrypts the specified AES encrypted text, represented as a base64 encoded string.
func Decrypt(encryptedValue string) (string, error) {
	bytes, err := base64.URLEncoding.DecodeString(encryptedValue)

	if err != nil || len(encryptedValue) < aes.BlockSize {
		return "", errors.New("Token is invalid.")
	}

	iv := bytes[:aes.BlockSize]
	cfb := cipher.NewCFBDecrypter(block, iv)

	bytes = bytes[aes.BlockSize:]
	cfb.XORKeyStream(bytes, bytes)

	value, err := base64.URLEncoding.DecodeString(string(bytes))

	if err != nil {
		return "", err
	}

	return string(value), nil
}
