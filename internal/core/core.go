// Global Variables, Models package
// project GophKeeper Yandex Practicum
// Created by Vasiliy Van-Zaam
package core

import (
	"fmt"
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
	Password(hash string, form *PasswordForm) error
	File(hash string, form *FileForm) error
	BankCard(hash string, form *BankCardForm) error
	Text(hash string, form *TextFomm) error
	TryPassword(hash string, count int) error
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
	DataTypePassword         DataType = "password"
	DataTypeCard             DataType = "card"
	DataTypeText             DataType = "text"
	DataTypeFile             DataType = "file"
	DataTypeUser             DataType = "user"
	DataTypeTryEnterPassword DataType = "trypassword"
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
	UserID   string `json:"user_id"`
	MetaData string `json:"metadata"`
	Password string `json:"value"`
	Login    string `json:"login"`    // login, email, phone
	Resource string `json:"resource"` // link site or application
}

// The bank card form for create a new bank card or display.
type BankCardForm struct {
	UserID     string `json:"user_id"`
	Metadata   string `json:"metadata"`
	Number     string `json:"number"`
	Date       string `json:"date"`
	CVC        string `json:"cvc"`
	ClientName string `json:"client_name"`
}
type TextFomm struct {
	MetaData string `json:"metadata"`
	Text     string `json:"text"`
}
type FileForm struct {
	MetaData string `json:"metadata"`
	Data     []byte `json:"text"`
	Name     string `json:"name"`
	Size     int64  `json:"size"`
}

func (m *ManagerData) ToManager() Manager {
	return NewManagerFromData(m)
}

type AppInfo struct {
	Version   string `json:"version"`
	SizeStore string `json:"size"`
	LastSync  string `json:"last_sync"`
}

type LocalStoreInfo interface {
	Size() int64
	LastSync() string
}
type ClientConfig interface {
	Version() string
}

func SizeToSize(size int64) string {
	toFloat := func(size int64, d float64) float64 {
		return float64(size) / d
	}
	sizeStr := ""
	switch {
	case len(fmt.Sprint(size)) > 9:
		sizeStr = fmt.Sprintf("%.3f %v", toFloat(size, 1000000000), "Gb")
	case len(fmt.Sprint(size)) > 6:
		sizeStr = fmt.Sprintf("%.3f %v", toFloat(size, 1000000), "Mb")
	case len(fmt.Sprint(size)) > 3:
		sizeStr = fmt.Sprintf("%.3f %v", toFloat(size, 1000), "kb")
	case len(fmt.Sprint(size)) <= 3:
		sizeStr = fmt.Sprint(size, "b")
	default:
	}
	return sizeStr
}

func NewAppInfo(clientConf ClientConfig, store LocalStoreInfo) *AppInfo {
	size := store.Size()
	sizeStr := SizeToSize(size)

	return &AppInfo{
		Version:   clientConf.Version(),
		SizeStore: sizeStr,
		LastSync:  store.LastSync(),
	}
}

type SyncInfo struct {
}
