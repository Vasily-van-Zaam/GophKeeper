// Global Variables, Models package
// project GophKeeper Yandex Practicum
// Created by Vasiliy Van-Zaam
package core

type Encryptor interface {
	Encrypt(hash string, data []byte) ([]byte, error)
	Decrypt(hash string, data []byte) ([]byte, error)
}

// someData.Set().MetaData("some info").DataType("pasword").Data(hash_password, []bayte(PasswordForm{})).
type Setter interface {
	MetaData(metaData string) Manager
	Data(hash string, data []byte) error
	Password(hash string, psw *PasswordForm) error
	BankCard(hash string, card *BankCardFomm) error
	Text(hash string, data string) error
}
type Getter interface {
	Data(hash string) ([]byte, error)
	InfoData() *InfoData
}

type Manager interface {
	Get() Getter
	Set() Setter
	AddEncription(encrypt Encryptor) Manager
}
