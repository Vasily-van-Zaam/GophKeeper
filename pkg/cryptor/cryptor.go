package cryptor

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

// Implements encrip interface!
type Encript interface {
	GetVersionHash(b []byte) (string, error)
	Encrypt(secret []byte, userData []byte) ([]byte, error)
	Decrypt(secret []byte, data []byte) ([]byte, error)
}

type criptor struct {
	aesgcm *cipher.AEAD
}

// Encrypt user key.
func (c *criptor) Encrypt(secret []byte, userData []byte) ([]byte, error) {
	hash := sha256.New()
	hash.Write(secret)
	k := hash.Sum(nil)
	aesblock, err := aes.NewCipher(k)
	if err != nil {
		return nil, err
	}
	aesgcm, err := cipher.NewGCM(aesblock)
	if err != nil {
		return nil, err
	}

	nonce, err := GenerateRandom(aesgcm.NonceSize())
	if err != nil {
		return nil, err
	}

	cryptedData := aesgcm.Seal(nil, nonce, userData, nil)

	// Здесь подмешиваем соль
	// пока подмешал в конец, но можно подмешать в любое место и при расшифровке вытащить соль
	cryptedData = append(cryptedData, nonce...)
	return cryptedData, nil
}

// Decrypt user key.
func (criptor) Decrypt(secret []byte, data []byte) ([]byte, error) {
	hash := sha256.New()
	hash.Write(secret)
	k := hash.Sum(nil)

	aesblock, err := aes.NewCipher(k)
	if err != nil {
		return nil, err
	}
	aesgcm, err := cipher.NewGCM(aesblock)
	if err != nil {
		return nil, err
	}
	nonce := data[len(data)-aesgcm.NonceSize():]
	decryptedData, err := aesgcm.Open(nil, nonce, data[:len(data)-aesgcm.NonceSize()], nil)
	if err != nil {
		return nil, err
	}
	return decryptedData, nil
}

// GetVersionHash implements Encript.
func (*criptor) GetVersionHash(b []byte) (string, error) {
	hash := sha512.New()
	hash.Write(b)
	res := fmt.Sprintf("%x", hash.Sum(nil))

	return string(res), nil
}

func New() Encript {
	return &criptor{}
}

// Generate Random.
func GenerateRandom(size int) ([]byte, error) {
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}
