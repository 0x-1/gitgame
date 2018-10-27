package CryptManager

import (
	"crypto/md5"
	"encoding/hex"
	"crypto/aes"
	"crypto/cipher"
	"io"
	"crypto/rand"
)
var plainkey string = "609132a0-fc3c-4b50-96d9-15efc3def63a"

//We hash to get aes min 32 character size key
func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func M_Encrypt(data []byte) ([]byte,error) {
	block, _ := aes.NewCipher([]byte(createHash(plainkey)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext, nil
}

func M_Decrypt(data []byte) ([]byte, error) {
	key := []byte(createHash(plainkey))
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}
