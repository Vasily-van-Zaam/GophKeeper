// Global Variables, Models package
// project GophKeeper Yandex Practicum
// Created by Vasiliy Van-Zaam
package core

import (
	"time"

	"github.com/google/uuid"
)

const (
	CtxVersionClientKey = "client_version"
	CtxTokenKey         = "token"
	CtxAcceptToken      = "acept_token"
)

type Encryptor interface {
	Encrypt(secret []byte, userData []byte) ([]byte, error)
	Decrypt(secret []byte, data []byte) ([]byte, error)
}

// someData.Set().MetaData("some info").DataType("pasword").Data(hash_password, []bayte(PasswordForm{})).
type Setter interface {
	AccessData(masterPsw string, user *User) error
	MetaData(metaData string) Manager
	Data(hash string, data []byte) error
	Password(hash string, psw *PasswordForm) error
	BankCard(hash string, card *BankCardFomm) error
	Text(hash string, data string) error
}

// Getter functions.
type Getter interface {
	Data(hash string) ([]byte, error)
	EncryptData() []byte
	InfoData() *InfoData
}

// Main Manager interface.
type Manager interface {
	Get() Getter
	Set() Setter
	AddEncription(encrypt Encryptor) Manager
	ToData() (*ManagerData, error)
}

// Main data types for saving.
type DataType string

// List data  types.
const (
	DataTypePassword DataType = "password"
	DataTypeCard     DataType = "card"
	DataTypeText     DataType = "text"
	DataTypeFile     DataType = "file"
	DataTypeUser     DataType = "user"
)

type DataGob struct {
	DataList []*ManagerData
}

// The ManagerData to save the database.
type ManagerData struct {
	Data []byte `json:"data"`
	InfoData
}

// The InfoData to save the database.
type InfoData struct {
	ID       *uuid.UUID `json:"id"`
	DataType string     `json:"data_type"`
	MetaData string     `json:"meta_data"`
	UserID   *uuid.UUID `json:"user_id"`
	Local    bool       `json:"local"`
	// /*
	// 	hash + updated_at is version changed
	// 	If the hash is not equal to the new hash,
	// 	then the version has changed.
	// 	The update date will be prompted to the client to choose which
	// 	version to keep in the central repository
	// */
	Hash      string     `json:"hash"`
	UpdatedAt *time.Time `json:"updated_at"`
	CreatedAt *time.Time `json:"created_at"`
}

// The password manager for create a new password or display.
type PasswordForm struct {
	Password string `json:"value"`
	Login    string `json:"login"`    // login, email, phone
	Resource string `json:"resource"` // link site or application
}

// The bank card form for create a new bank card or display.
type BankCardFomm struct {
	Number     string `json:"number"`
	Date       string `json:"date"`
	CVC        string `json:"cvc"`
	ClientName string `json:"client_name"`
}

func (m *ManagerData) ToManager() Manager {
	return NewManagerFromData(m)
}

type SyncInfo struct {
}
