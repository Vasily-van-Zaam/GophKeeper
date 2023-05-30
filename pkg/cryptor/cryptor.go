package cryptor

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"log"
)

// Implements encrip interface!
type Encript interface {
	GetVersionHash(b []byte) (string, error)
	Encrypt(hash string, data []byte) ([]byte, error)
	Decrypt(hash string, data []byte) ([]byte, error)
}

type criptor struct {
	aesgcm *cipher.AEAD
}

// Decrypt implements Encript.
func (c *criptor) Decrypt(hash string, data []byte) ([]byte, error) {
	res, err := (*c.aesgcm).Open(nil, []byte(hash), data, nil) // расшифровываем
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Encrypt implements Encript.
func (c *criptor) Encrypt(hash string, data []byte) ([]byte, error) {
	dst := (*c.aesgcm).Seal(nil, []byte(hash), data, nil)
	return dst, nil
}

// GetVersionHash implements Encript.
func (*criptor) GetVersionHash(b []byte) (string, error) {
	hash := sha512.New()
	hash.Write(b)
	res := fmt.Sprintf("%x", hash.Sum(nil))

	return string(res), nil
}

func New(key string) (Encript, error) {
	h := sha256.New()
	ky := []byte(key)

	h.Write(ky)

	k := h.Sum(nil)
	aesblock, err := aes.NewCipher(k)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(aesblock)
	if err != nil {
		return nil, err
	}

	sss := aesgcm.NonceSize()
	log.Println("????", sss)
	return &criptor{aesgcm: &aesgcm}, nil
}
