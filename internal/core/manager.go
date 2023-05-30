package core

import (
	"crypto/sha512"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func setVersionHash(b []byte) string {
	hash := sha512.New()
	hash.Write(b)
	res := fmt.Sprintf("%x", hash.Sum(nil))
	return res
}

type DataType string

const (
	DataTypePassword DataType = "password"
	DataTypeCard     DataType = "card"
	DataTypeText     DataType = "text"
	DataTypeFile     DataType = "file"
)

// The InfoData for saving.
type InfoData struct {
	ID       *uuid.UUID `json:"id"`
	DataType string     `json:"data_type"`
	MetaData string     `json:"meta_data"`
	UserID   *uuid.UUID `json:"user_id"`
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

// The data manager model for saving.
type manager struct {
	infoData  *InfoData
	data      *[]byte // encrypted
	setter    Setter
	getter    Getter
	encryptor Encryptor
}

type setter struct {
	data *manager
}

// BankCard implements Setter.
func (s *setter) BankCard(hash string, card *BankCardFomm) error {
	if s.data.encryptor == nil {
		return errors.New("need add encrypter")
	}
	d, err := json.Marshal(card)
	if err != nil {
		return err
	}
	bs, err := s.data.encryptor.Encrypt(hash, d)
	if err != nil {
		return err
	}
	s.data.infoData.Hash = setVersionHash(d)
	s.data.infoData.DataType = string(DataTypeCard)
	s.data.data = &bs
	return nil
}

// Password implements Setter.
func (s *setter) Password(hash string, psw *PasswordForm) error {
	if s.data.encryptor == nil {
		return errors.New("need add encrypter")
	}
	d, err := json.Marshal(psw)
	if err != nil {
		return err
	}
	bs, err := s.data.encryptor.Encrypt(hash, d)
	if err != nil {
		return err
	}
	s.data.infoData.Hash = setVersionHash(d)
	s.data.infoData.DataType = string(DataTypePassword)
	s.data.data = &bs
	return nil
}

// Text implements Setter.
func (s *setter) Text(hash string, text string) error {
	if s.data.encryptor == nil {
		return errors.New("need add encrypter")
	}
	d := []byte(text)
	bs, err := s.data.encryptor.Encrypt(hash, d)
	if err != nil {
		return err
	}
	s.data.infoData.Hash = setVersionHash(d)
	s.data.infoData.DataType = string(DataTypeText)
	s.data.data = &bs
	return nil
}

// Data implements Setter.
func (s *setter) Data(hash string, data []byte) error {
	if s.data.encryptor == nil {
		return errors.New("need add encrypter")
	}
	bs, err := s.data.encryptor.Encrypt(hash, data)
	if err != nil {
		return err
	}
	s.data.infoData.Hash = setVersionHash(data)
	s.data.infoData.DataType = string(DataTypeFile)
	s.data.data = &bs
	return nil
}

// MetaData implements Setter.
func (s *setter) MetaData(metaData string) Manager {
	s.data.infoData.MetaData = metaData
	return s.data
}

type getter struct {
	data *manager
}

// Data implements Getter.
func (g *getter) Data(hash string) ([]byte, error) {
	if g.data.encryptor == nil {
		return nil, errors.New("need add encrypter")
	}
	bs, err := g.data.encryptor.Decrypt(hash, *g.data.data)
	if err != nil {
		return nil, err
	}
	return bs, nil
}

// InfoData implements Getter.
func (g *getter) InfoData() *InfoData {
	return g.data.infoData
}

// AddEncription implements Manager.
func (d *manager) AddEncription(encrypt Encryptor) Manager {
	d.encryptor = encrypt
	return d
}

// Get implements Manager.
func (d *manager) Get() Getter {
	return d.getter
}

// Set implements Manager.
func (d *manager) Set() Setter {
	return d.setter
}

// Create new manager for store.
func NewData() Manager {
	data := &manager{
		infoData: &InfoData{},
	}
	data.getter = &getter{
		data: data,
	}
	data.setter = &setter{
		data: data,
	}
	return data
}

// Create manager form storage.
func NewDataFrom(info *InfoData, data []byte) Manager {
	d := &manager{
		infoData: info,
		data:     &data,
	}
	d.getter = &getter{
		data: d,
	}
	d.setter = &setter{
		data: d,
	}
	d.data = &data
	return d
}
