package cryptor

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
)

// Implements encript interface!
type Encryptor interface {
	Encrypt(secret []byte, userData []byte) ([]byte, error)
	Decrypt(secret []byte, data []byte) ([]byte, error)
}

type cryptor struct{}

// Encrypt user key.
func (c *cryptor) Encrypt(secret []byte, userData []byte) ([]byte, error) {
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

	nonce, err := generateRandom(aesgcm.NonceSize())
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
func (cryptor) Decrypt(secret []byte, data []byte) ([]byte, error) {
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

// Create new encryption.
func New() Encryptor {
	return &cryptor{}
}

// Generate Random.
func generateRandom(size int) ([]byte, error) {
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}
